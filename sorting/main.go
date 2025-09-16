package main

import (
	"fmt"
	"time"
)

/*
Look through the whole array and find the minimum element.
Swap that minimum with the element at the current starting position.
Now the first position is “locked” (sorted).
Repeat for the rest of the array, moving the boundary one step forward each time.
Keep going until the whole array is sorted.
*/

func SelectionSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if minIdx != i {
			arr[minIdx], arr[i] = arr[i], arr[minIdx]
		}
	}
	return arr
}

/*
Compare each pair of adjacent elements.
If the left one is bigger than the right, swap them.
After one full pass, the largest element “bubbles” to the end.
Ignore the last element (it’s in place) and repeat with the remaining unsorted part.
Keep doing this until no swaps are needed.
*/
func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func RecursionBubbleSort(arr []int, end int) []int {

	if end == 1 {
		return arr
	}

	for i := range end {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
	return RecursionBubbleSort(arr, end-1)
}

/*
Insertion Sort
Builds a sorted part of the array one element at a time.
For each i, take arr[i] and move it left until it’s in the right spot.
Best case (already sorted): O(n). Worst/Average: O(n²).
Think: sorting playing cards in your hand.
*/
func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if arr[j-1] > arr[j] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}

func RecursionInsertionSort(arr []int, n int) {
	if n <= 1 {
		return
	}

	RecursionInsertionSort(arr, n-1)

}

/*
Merge Sort
Divide-and-conquer sorting, O(n log n), stable.
Steps:
1. Split array into two halves (left, right).
2. Recursively sort each half.
3. Merge halves back in sorted order.
Key: main loop merges until one half ends, then append leftovers.
O(n log n) time and uses O(n) extra space.
*/

func Merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func MergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return Merge(left, right)
}

func Partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func QuickSort(arr []int, low, high int) {
	if low < high {
		pi := Partition(arr, low, high)

		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func main() {
	start := time.Now()
	arr := []int{1, 4, 5, 3, 8, 0, 2, 6}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	fmt.Println(time.Since(start))
}
