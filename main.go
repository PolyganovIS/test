package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	reversed := 0
	original := x

	for x > 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}
	return original == reversed
}

func main() {
	// Проверяем несколько чисел на палиндромность
	nums := []int{121, -121, 10, 12321, 123321}

	for _, num := range nums {
		fmt.Printf("Число %d является палиндромом? %t\n", num, isPalindrome(num))
	}
}
