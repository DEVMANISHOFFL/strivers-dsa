package main

import (
	"fmt"
	"sort"
)

func LargestArray(arr []int) int {
	sort.Ints(arr)
	return arr[0]
}

func SecondLargest(arr []int) int {
	if len(arr) < 2 {
		panic("more than two elements required")
	}

	largest, sLargest := arr[0], -1<<31
	for _, val := range arr[1:] {
		if val > largest {
			sLargest = largest
			largest = val
		} else if val > sLargest && val != largest {
			sLargest = val
		}
	}
	return sLargest

}

func CheckIfSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func removeDuplicates(nums []int) (int, []int) {
	newArr := []int{nums[0]}
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			newArr = append(newArr, nums[i])
			count++
		}
	}
	return count, newArr
}

func countDuplicatedAndRemove(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	i := 0
	for j := 1; j < len(arr); j++ {
		if arr[j] != arr[i] {
			i++
			arr[i] = arr[j]
		}
	}
	return i + 1
}

func LeftRotate(arr []int) []int {
	n := len(arr)
	pivot := arr[0]
	for i := 0; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[n-1] = pivot
	return arr
}

func RighRotate(arr []int) []int {
	n := len(arr)
	pivot := arr[n-1]
	for i := n - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = pivot
	return arr
}

func LeftRotateByDPlace(arr []int, d int) []int {
	n := len(arr)
	d = d % n

	if n == 0 {
		return arr
	}

	if d == 0 {
		return arr
	}

	newArr := []int{}

	for i := 0; i < d; i++ {
		newArr = append(newArr, arr[i])
	}

	for i := 0; i < n-d; i++ {
		arr[i] = arr[i+d]
	}

	arr = append(arr[:n-d], newArr...)
	return arr
}

func LeftRotateByDPlaceByReversing(arr []int, d int) []int {
	n := len(arr)
	d = d % n

	for i, j := 0, d; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for i, j := d+1, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func NumToLast(arr []int, num int) []int {
	n := len(arr)
	notNum := []int{}
	count := 0
	for i := 0; i < n; i++ {
		if arr[i] == num {
			count++
		}
		if arr[i] != num {
			notNum = append(notNum, arr[i])
		}
	}
	for i := 0; i < count; i++ {
		notNum = append(notNum, num)
	}

	return notNum
}

func NumToLastOptimal(arr []int, num int) []int {
	n := len(arr)
	j := -1

	for i := 0; i < n; i++ {
		if arr[i] == num {
			j = i
			break
		}
	}

	if j == -1 {
		return arr
	}

	for i := j + 1; i < n; i++ {

		if arr[i] != num {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	return arr
}

func LinearSearch(arr []int, num int) {
	// var found int
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			fmt.Println(i)
			break
		}
	}
}

func UnionOfTwoArray(arr1, arr2 []int) []int {
	duplicates := []int{}

	union := []int{}

	for i := 0; i < len(arr1); i++ {
		duplicates = append(duplicates, arr1[i])
	}
	for i := 0; i < len(arr2); i++ {
		duplicates = append(duplicates, arr2[i])
	}

	sort.Ints(duplicates)
	if len(duplicates) == 0 {
		return []int{}
	}

	for i := 1; i < len(duplicates); i++ {
		if duplicates[i] != duplicates[i-1] {
			union = append(union, duplicates[i])
		}
	}
	return union
}

func Intersection(arr1, arr2 []int) []int {
	n1, n2 := len(arr1), len(arr2)
	res := []int{}
	i, j := 0, 0
	for i < n1 && j < n2 {
		if arr1[i] < arr2[j] {

			i++
		} else if arr1[i] > arr2[j] {
			j++
		} else {
			res = append(res, arr1[i])
			i++
			j++
		}
	}
	return res
}

func FindMissingElement(arr []int, till int) []int {
	n := len(arr)
	notAvail := []int{}
	for i := 0; i < till; i++ {
		found := false
		for j := 0; j < n; j++ {
			if arr[j] == i {
				found = true
				break
			}
		}
		if !found {
			notAvail = append(notAvail, i)
		}
	}
	return notAvail
}

func FindMissingElementSum(arr []int) int {
	n := len(arr)
	var missing int
	var sum int
	actualSum := n * (n + 1) / 2
	for i := 0; i < n; i++ {
		sum = sum + arr[i]
	}
	missing = actualSum - sum
	return missing
}

func CountConsecutive(arr []int) int {
	var max int
	var count int
	for _, v := range arr {
		if v == 1 {
			count++
			if count > max {
				max = count
			}
		} else {
			count = 0
		}

	}
	return max
}

func OccuredOnce(arr []int) int {
	sort.Ints(arr)
	fmt.Println(arr)
	var once int
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			once = arr[i]
		}
	}
	return once
}

func maxSubarray(arr []int, k int) ([]int, int) {
	n := len(arr)
	res := []int{}
	maxLen := 0
	for i := range arr {
		sum := 0
		for j := i; j < n; j++ {
			sum = sum + arr[j]
			subLen := j - i + 1
			if sum == k && subLen > maxLen {
				maxLen = subLen
				res = arr[i : j+1]
			}
		}
	}
	return res, len(res)
}

func prefixSum(arr []int) []int {
	pSum := make([]int, len(arr)+1)
	sum := 0
	for i := 1; i <= len(arr); i++ {
		sum = sum + arr[i]
		pSum[i] = pSum[i-1] + arr[i-1]
	}
	return pSum

}

func main() {

	// arr := []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 3, 4, 6, 7}
	arr := []int{1, 2, 3, 4, 5}
	// arr := []int{0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1, 1}

	fmt.Println(prefixSum(arr))
}
