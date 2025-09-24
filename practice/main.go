package main

import "fmt"

func prac(arr []int, tar int) [][]int {
	seen := make(map[int]int)
	pairs := [][]int{}

	for i := 0; i < len(arr); i++ {
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
func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// arr := []int{1, 1, 2, 0, 4, 3, 4, 4, 4, 4, 5, 4, 4, 1, 4, 4, 4, 4, 4, 4}
	fmt.Println(prac(arr, 5))

}
