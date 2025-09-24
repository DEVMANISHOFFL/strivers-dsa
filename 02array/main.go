package main

import (
	"fmt"
	"sort"
)

func twoSum(arr []int, tar int) bool {
	n := len(arr)
	isPossible := false
	for i := 0; i < n; i++ {
		for j := 0; j < n && j != i; j++ {
			sum := arr[i] + arr[j]
			if sum == tar {
				isPossible = true
			}
		}
	}
	return isPossible
}

func twoSumPairs(arr []int, tar int) [][]int {
	n := len(arr)
	pairs := [][]int{}

	for i := 0; i < n; i++ {
		for j := 0; j < n && j != i; j++ {
			if arr[i]+arr[j] == tar {
				pairs = append(pairs, []int{arr[i], arr[j]})
			}
		}
	}
	return pairs
}

func twoSumPairsSorting(nums []int, target int) [][]int {
	sort.Ints(nums) // O(n log n)
	l, r := 0, len(nums)-1
	pairs := [][]int{}

	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			pairs = append(pairs, []int{nums[l], nums[r]})
			l++
			r--
			for l < r && nums[l] == nums[l-1] {
				l++
			}
			for l < r && nums[r] == nums[r+1] {
				r--
			}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return pairs
}
func TwoSumHashing(arr []int, tar int) [][]int {
	pairs := [][]int{}
	seen := make(map[int]int)
	n := len(arr)
	for i := 0; i < n; i++ {
		complement := tar - arr[i]
		if count, ok := seen[complement]; ok {
			for range count {

				pairs = append(pairs, []int{complement, arr[i]})
			}
		}
		seen[arr[i]]++
	}
	return pairs
}

func DutchNationalFlag(nums []int) {
	low, mid, high := 0, 0, len(nums)-1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[mid], nums[low] = nums[low], nums[mid]
			low++
			mid++
		case 1:
			mid++
		default:
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		}
	}
}

func MajorityElement(arr []int) int {
	n := len(arr)
	seen := make(map[int]int)

	for i := range n {
		seen[arr[i]]++
		if seen[arr[i]] > n/2 {
			return arr[i]
		}
	}
	return -1
}

func MajorityElementOptimal(arr []int) int {
	n := len(arr)
	var count, element int

	for i := 0; i < n; i++ {
		if count == 0 {
			element = arr[i]
			count = 1
		} else if arr[i] == element {
			count++
		} else {
			count--
		}
		count2 := 0
		for i := 0; i < n; i++ {
			if arr[i] == element {
				count2++
			}
			if count2 > n/2 {
				return element
			}
		}
	}
	return -1
}

func maxSubArraySum(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}

	maxSum := arr[0]

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func kadane(arr []int) int {
	maxSoFar := arr[0]
	curr := arr[0]

	for i := 1; i < len(arr); i++ {
		sum := curr + arr[i]

		curr = max(sum, arr[i])
		if curr > maxSoFar {

			maxSoFar = curr
		}
	}
	return maxSoFar
}

func kadanes(arr []int) (int, []int) {
	n := len(arr)
	maxSoFar, curr := arr[0], arr[0]
	subArray := []int{}

	if n == 0 {
		return 0, nil
	}

	for i := 1; i < n; i++ {
		curr = max(arr[i], arr[i]+curr)
		maxSoFar = max(curr, maxSoFar)
	}
	return maxSoFar, subArray

}

func maxS(arr []int) ([]int, int) {
	n := len(arr)
	if n == 0 {
		return nil, 0
	}

	start, end, tempStart := 0, 0, 0
	maxSoFar, curr := arr[0], arr[0]

	for i := 1; i < n; i++ {
		if curr+arr[i] < arr[i] {
			curr = arr[i]
			tempStart = i
		} else {
			curr += arr[i]

		}
		if curr > maxSoFar {
			maxSoFar = curr
			start = tempStart
			end = i
		}
	}
	sub := make([]int, end-start+1)
	copy(sub, arr[start:end+1])
	return sub, maxSoFar
}

func BuyAndSell(arr []int) int {
	n := len(arr)
	maxProfit := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			profit := arr[j] - arr[i]
			if maxProfit < profit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

func BuyAndSellOptimal(arr []int) int {
	n := len(arr)

	maxProfit := 0
	minPrice := arr[0]

	for i := 1; i < n; i++ {
		cost := arr[i] - minPrice
		maxProfit = max(maxProfit, cost)

		minPrice = min(arr[i], minPrice)
	}
	return maxProfit
}

func main() {
	// arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 34, 56, 78, 91, 23, 45, 67, 89, 12, 34, 56, 78, 90}
	// arr := []int{0, 1, 1, 1, 1, 0, 0, 0, 2, 2, 2, 2, 2, 2, 1, 0, 1, 2}
	// arr := []int{7, 1, 5, 3, 6, 4}
	arr := []int{-2, 1, -3, 4, -1, 2, 1, 5, 4}

	// fmt.Println(MajorityElement(arr))
	// fmt.Println(kadanes(arr))
	// fmt.Println(kadane(arr))
	fmt.Println(BuyAndSell(arr))

}
