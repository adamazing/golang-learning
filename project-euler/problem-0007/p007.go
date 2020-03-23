package main

import (
	"fmt"
	"time"

	"../lib/primesieve"
)

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the
// 6th prime is 13.
//
// What is the 10 001st prime number?
func main() {
	var result int
	startTime := time.Now()
	fmt.Println("Problem 7")

	result = primesieve.NthPrime(10001)

	fmt.Printf("Answer: %d\n", result)

	fmt.Printf("Took: %v", time.Now().Sub(startTime))
}
