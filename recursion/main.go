package main

import "fmt"

var count int

func f() {
	for count <= 100000 {
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

func main() {
	printNToOneLinearOneParam(10)
}
