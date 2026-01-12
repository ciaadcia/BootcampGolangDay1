package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("No 1 - Find Divisor")
	findDivisor(6)
	findDivisor(24)
	findDivisor(7)

	fmt.Println("\nNo 2 - Extract Digit")
	extractDigit(12234)
	extractDigit(5432)
	extractDigit(1278)

	fmt.Println("\nNo 3 - Triangle Star")
	triangleStar(5)

	fmt.Println("\nNo 4 - Number Pyramid")
	numberPyramidInput()

	fmt.Println("\nNo 8 - Palindrome String")
	fmt.Println(isPalindrome("Kasur ini rusak"))
	fmt.Println(isPalindrome("tamaT"))
	fmt.Println(isPalindrome("Aku Usa"))

	fmt.Println("\nNo 9 - Reverse String")
	fmt.Println(reverse("ABCD"))
	fmt.Println(reverse("tamaT"))
	fmt.Println(reverse("XYnb"))

	fmt.Println("\nNo 10 - Check Braces")
	fmt.Println(checkBraces("()"))
	fmt.Println(checkBraces("()()"))
	fmt.Println(checkBraces("(()"))
	fmt.Println(checkBraces("()(()"))

	fmt.Println("\nNo 11 - Palindrome Number")
	fmt.Println(isNumberPalindrome(121))
	fmt.Println(isNumberPalindrome(2147447412))
	fmt.Println(isNumberPalindrome(333))
	fmt.Println(isNumberPalindrome(110))
	fmt.Println(isNumberPalindrome(11))
}

/*
No 1
*/
func findDivisor(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

/*
No 2
*/
func extractDigit(n int) {
	for n > 0 {
		fmt.Print(n%10, " ")
		n /= 10
	}
	fmt.Println()
}

/*
No 3
*/
func triangleStar(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= i {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

/*
No 4
*/
func numberPyramidInput() {
	var n int
	fmt.Print("Masukkan jumlah baris piramida: ")
	fmt.Scanln(&n)

	for i := n; i >= 1; i-- {

		for s := 0; s < n-i; s++ {
			fmt.Print("  ")
		}

		for j := i; j >= 1; j-- {
			fmt.Print(j, " ")
		}

		for j := 2; j <= i; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
}

/*
No 8
*/
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	return s == reverse(s)
}

/*
No 9
*/
func reverse(s string) string {
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}

/*
No 10
*/
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

/*
No 11
*/
func isNumberPalindrome(n int) bool {
	ori := n
	rev := 0

	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return ori == rev
}
