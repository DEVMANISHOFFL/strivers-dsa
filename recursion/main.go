package main

import "fmt"

var count int

func f() {
	for count <= 100 {
		fmt.Println(count)
		count++
		f()
	}
}

func printNameTwoParam(i, n int) {
	if i > n {
		return
	}
	fmt.Println("Recursion Two Params")
	printNameTwoParam(i+1, n)
}

func printNameOneParam(n int) {
	if n == 0 {
		return
	}
	fmt.Println("Recursion One Param")
	printNameOneParam(n - 1)
}

func printNToOneLinearOneParam(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNToOneLinearOneParam(n - 1)
}

func printNtoOneLinearTwoParam(i, n int) {
	if i < 1 {
		return
	}
	fmt.Println(i)
	printNtoOneLinearTwoParam(i-1, n)
}

func printOneToNLinearTwoParam(i, n int) {
	if i > n {
		return
	}
	fmt.Println(i)
	printOneToNLinearTwoParam(i+1, n)
}

func printOneToNLinearOneParam(n int) {
	if n < 1 {
		return
	}
	printOneToNLinearOneParam(n - 1)
	fmt.Println(n)
}

func sumOfFristN(sum, n int) int {
	if n == 0 {
		return sum
	}
	return sumOfFristN(sum+n, n-1)
}

func sumOfFristNIdiomatic(n int) int {
	if n == 0 {
		return 0
	}
	return n + sumOfFristNIdiomatic(n-1)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func reverseArray(a []int) []int {
	l, r := 0, len(a)-1
	if len(a) <= 1 {
		return a
	}
	a[l], a[r] = a[r], a[l]
	reverseArray(a[1 : len(a)-1])
	return a
}

func reverseArrayWithHelper(a []int) []int {
	return helperReverse(a, 0, len(a)-1)
}

func helperReverse(a []int, l, r int) []int {
	if l >= r {
		return a
	}

	a[l], a[r] = a[r], a[l]
	return helperReverse(a, l+1, r-1)
}

func isPalindromArray(a []int) bool {
	l, r := 0, len(a)-1
	if len(a) <= 1 {
		return true
	}
	if a[l] != a[r] {
		return false
	}
	return isPalindromArray(a[1 : len(a)-1])
}

func isPalindromString(s string) bool {
	runes := []rune(s)
	if len(runes) <= 1 {
		return true
	}

	if runes[0] != runes[len(runes)-1] {
		return false
	}
	return isPalindromString(string(runes[1 : len(runes)-1]))
}
func fiboNum(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fiboNum(n-2) + fiboNum(n-1)

}

func main() {
	// a := []int{1, 2, 3, 2, 1, 1, 2, 3, 2, 1}

	fmt.Println(fiboNum(40))
}
