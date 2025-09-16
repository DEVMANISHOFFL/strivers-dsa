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

func main() {

	arr := []int{3, 4, 5, 5, 5, 5, 1, 1, 1, 2, 6, 6, 6, 1}
	fmt.Println(removeDuplicates(arr))
}
