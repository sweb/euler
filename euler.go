package main

import (
	"fmt"
	"github.com/sweb/euler/problems"
)

func main() {
	multiplesCap := 1000
	fmt.Printf("Result of MultiplesOf3And5(%v) is %v\n", multiplesCap, problems.MultiplesOf3And5(multiplesCap))

	lb, ub := 100, 999
	fmt.Printf("Result of LargestPalindromeProduct(%v, %v) is %v\n", lb, ub, problems.LargestPalindromeProduct(lb, ub))
}
