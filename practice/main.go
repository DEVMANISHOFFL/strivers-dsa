package main

import "fmt"

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
	arr := []int{1, 3, 2, 5, 4, 7, 6, 9, 8, 0, 0}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
