package main

import "fmt"

func recursion(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	return recursion(n-1) + recursion(n-2)

}

func main() {
	// fmt.Println(recursion([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(recursion(20))
}
