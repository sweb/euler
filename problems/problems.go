package problems

import (
	"strconv"
)

func MultiplesOf3And5(cap int) int {
	var sum int = 0

	for i := 0; i < cap; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func LargestPalindromeProduct(lowBound int, highBound int) int {
	product := 0
	for i := lowBound; i <= highBound; i++ {
		for j := i; j <= highBound; j++ {
			currentProduct := i * j
			if currentProduct > product && isPalindrome(currentProduct) {
				product = currentProduct
			}
		}
	}
	return product
}

func isPalindrome(number int) bool {
	tmp := strconv.Itoa(number)
	return tmp == reverseString(tmp)
}

func reverseString(s string) string {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		oldChar := runes[i]
		runes[i] = runes[length-1-i]
		runes[length-1-i] = oldChar
	}
	return string(runes)
}
