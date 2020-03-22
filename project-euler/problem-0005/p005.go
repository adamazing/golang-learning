package main

import (
	"fmt"
	"time"
)

// 2520 is the smallest number that can be divided by each of the numbers from
// 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?

func maxPossibleResult(divisors []int) int {
	result := 1
	for _, v := range divisors {
		result = result * v
	}
	return result
}

func isDivisibleByAll(num int, divisors []int) bool {
	for k := 0; k < len(divisors); k++ {
		if num%divisors[k] != 0 {
			return false
		}
	}
	return true
}

func smallestPositiveNumberDivisibleByNumbersLessThan(n int) int {
	var valsUnderLimit = make([]int, n)
	for i := range valsUnderLimit {
		valsUnderLimit[i] = i + 1
	}

	for j := n; j <= maxPossibleResult(valsUnderLimit); j += n {
		if isDivisibleByAll(j, valsUnderLimit) {
			return j
		}
	}

	fmt.Println(valsUnderLimit)
	return 0
}

func main() {
	var result int
	startTime := time.Now()
	fmt.Println("Problem 5")

	result = smallestPositiveNumberDivisibleByNumbersLessThan(20)

	fmt.Printf("Answer: %d\n", result)

	fmt.Printf("Took: %v\n", time.Now().Sub(startTime))
}
