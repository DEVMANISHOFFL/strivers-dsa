package main

import (
	"fmt"
)

func pattern1(n int) {
	for range n {
		for range n {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func pattern2(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("+ ")
		}
		fmt.Println("")
	}
}

func pattern3(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func pattern4(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func pattern5(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print("+ ")
		}
		fmt.Println()
	}
}

func pattern6(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i+1; j++ {
			fmt.Print(j)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func pattern7(n int) {
	for i := 1; i <= n; i++ {
		for k := 0; k < n-i; k++ {
			fmt.Print("   ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Print(" + ")
		}

		fmt.Println()
	}
}

func pattern8(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for k := 2*(n-i) - 1; k > 0; k-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func pattern9(n int) {
	for i := range n {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= 2*i; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*n-2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func pattern10(n int) {
	for i := 1; i <= 2*n-1; i++ {
		var stars int
		if i <= n {
			stars = i
		} else {
			stars = 2*n - i
		}
		for j := 0; j < stars; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func pattern11(n int) {
	for i := 0; i < n; i++ {
		var start int
		if i%2 == 0 {
			start = 1
		} else {
			start = 0
		}

		for j := 0; j <= i; j++ {
			fmt.Print(start)
			start = 1 - start
		}
		fmt.Println()
	}
}

func pattern12(n int) {
	for i := 1; i < n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		for j := 0; j < 2*(n-i-1); j++ {
			fmt.Print(" ")
		}
		for j := i; j >= 1; j-- {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func pattern13(n int) {
	var start int
	start = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(" %d ", start)
			start++
		}
		fmt.Println()
	}
}

func pattern14(n int) {
	for i := 0; i < n; i++ {
		for ch := 'A'; ch <= 'A'+rune(i); ch++ {
			fmt.Printf("%c", ch)
		}
		fmt.Println()
	}
}

func pattern15(n int) {
	for i := 0; i < n; i++ {
		for ch := 'A'; ch < 'A'+rune(n-i); ch++ {
			fmt.Printf("%c", ch)
		}
		fmt.Println()
	}
}

func pattern16(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			ch := 'A' + rune(i-1)
			fmt.Printf("%c", ch)
		}
		fmt.Println()
	}
}

func pattern17(n int) {
	for i := 1; i <= n; i++ {
		for s := 0; s < n-i; s++ {
			fmt.Print(" ")
		}
		peak := 'A' + rune(i-1)
		for ch := 'A'; ch <= peak; ch++ {
			fmt.Printf("%c", ch)
		}

		for ch := peak - 1; ch >= 'A'; ch-- {
			fmt.Printf("%c", ch)
		}
		fmt.Println()
	}
}

func pattern18(n int) {
	end := 'A' + rune(n)
	for i := 1; i <= n; i++ {
		start := end - rune(i)
		for ch := start; ch < end; ch++ {
			fmt.Printf(" %c ", ch)
		}
		fmt.Println()
	}
}

func pattern19(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		for j := 1; j <= 2*i; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		for j := 0; j < 2*(n-i); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func pattern20(n int) {
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		for j := 0; j < 2*(n-i); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 1; i <= n-1; i++ {
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		for j := 0; j < 2*(i); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < n-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func pattern21(n int) {
	for i := range n {
		for j := range n {
			if i == 0 || j == 0 || i == n-1 || j == n-1 {

				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}

		}
		fmt.Println()
	}
}

func main() {
	pattern21(5)

}
