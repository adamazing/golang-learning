package main

import (
	"fmt"
	"time"
)

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the
// 6th prime is 13.
//
// What is the 10 001st prime number?

// NthPrime function
//   n            The upper limit.
//   lowerLimit   The lower limit.
func nthPrime(n int) int {
	counter := 1
	var isPrime bool
	var j int
	var primes []int
	primes = append(primes, 2)
	for len(primes) < n {
		counter += 2
		j = 0
		isPrime = true
		for j < len(primes) && primes[j]^2 <= counter {
			if counter%primes[j] == 0 {
				isPrime = false
				break
			}
			j++
		}
		if isPrime {
			primes = append(primes, counter)
		}
	}
	return primes[len(primes)-1]
}

func main() {
	var result int
	startTime := time.Now()
	fmt.Println("Problem 7")

	result = nthPrime(10001)

	fmt.Printf("Answer: %d\n", result)

	fmt.Printf("Took: %v", time.Now().Sub(startTime))
}
