package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func DigitExtraction(n int) {
	var count int
	for n > 0 {
		digit := n % 10
		fmt.Println(digit)
		n /= 10
		count++
	}
	fmt.Printf("count: %d", count)
	fmt.Println()
}

func ReverseDigits(n int) int {
	var reverse int
	for n > 0 {
		lastDigit := n % 10
		n /= 10
		reverse = reverse*10 + lastDigit
	}
	return reverse
}

func CheckPalindrome(n int) bool {
	var rev int
	num := n

	for num > 0 {
		lastDigit := num % 10
		num /= 10
		rev = rev*10 + lastDigit
	}
	if rev == n {
		return true
	} else {
		return false
	}
}

func CheckPalindromeIdomatic(n int) bool {
	rev, num := 0, n
	for num > 0 {
		rev = rev*10 + num%10
		num /= 10
	}
	return rev == n
}
func ArmstrongNum(n int) bool {
	num, sum, count := n, 0, 0

	for num > 0 {
		count++
		num /= 10
	}

	num = n
	for num > 0 {
		digits := num % 10
		sum += int(math.Pow(float64(digits), float64(count)))
		num /= 10
	}
	return sum == n
}

func PrintAllDivisors(n int) {
	start := time.Now()

	num := n
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println(n)
	elasped := time.Since(start)
	fmt.Println(elasped)

}

func PrintAllDivisorsFaster(n int) {
	start := time.Now()
	limit := int(math.Sqrt(float64(n)))
	divs := make([]int, 0, 2*limit)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divs = append(divs, i)
			if i != n/i { // to not to repeat that perfect square line 6 and 6 in case of n = 36
				divs = append(divs, n/i)
			}
		}
	}
	sort.Ints(divs)

	for _, d := range divs {
		fmt.Println(d)
	}

	elapsed := time.Since(start)
	fmt.Println("Time taken:", elapsed)
}

func IsPrimeNum(n int) bool {
	count := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			count++

			if i != n%i {
				count++
			}

		}

	}
	return count == 2
}

func main() {
	fmt.Println(IsPrimeNum(13))
}
