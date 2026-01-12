package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("No 1")
	findDivisor(6)
	findDivisor(24)
	findDivisor(7)

	fmt.Println("No 2")
	extractDigit(12234)
	extractDigit(5432)
	extractDigit(1278)

	fmt.Println("No 3")
	triangleStarLeft(5)
	fmt.Println()
	triangleStarRight(5)

	fmt.Println("No 4")
	numberPyramidInput()

	fmt.Println("No 5")
	patternFive(5)

	fmt.Println("No 6")
	patternSix(5)

	fmt.Println("No 8")
	fmt.Println(isPalindrome("Kasur ini rusak"))
	fmt.Println(isPalindrome("tamaT"))
	fmt.Println(isPalindrome("Aku Usa"))

	fmt.Println("No 9")
	fmt.Println(reverse("ABCD"))
	fmt.Println(reverse("tamaT"))
	fmt.Println(reverse("XYnb"))

	fmt.Println("No 10")
	fmt.Println(checkBraces("()"))
	fmt.Println(checkBraces("()()"))
	fmt.Println(checkBraces("(()"))
	fmt.Println(checkBraces("()(()"))

	fmt.Println("No 11")
	fmt.Println(isNumberPalindrome(121))
	fmt.Println(isNumberPalindrome(2147447412))
	fmt.Println(isNumberPalindrome(333))
	fmt.Println(isNumberPalindrome(110))
	fmt.Println(isNumberPalindrome(11))
}

func findDivisor(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func extractDigit(n int) {
	for n > 0 {
		fmt.Print(n%10, " ")
		n /= 10
	}
	fmt.Println()
}

func triangleStarLeft(n int) {
	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func triangleStarRight(n int) {
	for i := 1; i <= n; i++ {
		for s := 1; s <= n-i; s++ {
			fmt.Print("  ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func numberPyramidInput() {
	var n int
	fmt.Print("Masukkan jumlah baris piramida: ")
	fmt.Scanln(&n)

	for i := n; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			fmt.Print(j, " ")
		}
		for j := 2; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}

func patternFive(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j%2 == 1 {
				fmt.Print(i, " ")
			} else {
				fmt.Print(n-i+1, " ")
			}
		}
		fmt.Println()
	}
}

func patternSix(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("- ")
			} else {
				fmt.Print(j, " ")
			}
		}
		fmt.Println()
	}
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	return s == reverse(s)
}

func reverse(s string) string {
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}

func checkBraces(s string) bool {
	count := 0
	for _, c := range s {
		if c == '(' {
			count++
		} else if c == ')' {
			count--
			if count < 0 {
				return false
			}
		}
	}
	return count == 0
}

func isNumberPalindrome(n int) bool {
	ori := n
	rev := 0
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return ori == rev
}
