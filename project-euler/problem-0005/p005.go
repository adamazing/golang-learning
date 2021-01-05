package main

import (
	"fmt"
	"math"
	"math/big"
	"time"
)

// 2520 is the smallest number that can be divided by each of the numbers from
// 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?


// Sieve function
//   n            The upper limit.
//   lowerLimit   The lower limit.
func sieve(n, lowerLimit int) []int {
	var prime = make([]bool, n+1)
	var primes []int
	// fill it with true values
	for i := range prime {
		prime[i] = true
	}
	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		if big.NewInt(int64(i)).ProbablyPrime(10) {
			for j := i * i; j <= n; j += i {
				prime[j] = false
			}
		}
	}
	for i, v := range prime {
		if v && i > lowerLimit {
			primes = append(primes, i)
		}
	}
	return primes
}

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

func refinedSmallestPositiveNumberDivisibleByNumbersLessThan(n int) int {
	sieveValues := sieve(n, 0)
	result := 1
	for i := 0; i < len(sieveValues); i++ {
		var a float64 = math.Floor(float64(math.Log(float64(n)) / math.Log(float64(sieveValues[i]))))
		result = result * (int(math.Pow(float64(sieveValues[i]), a)))
	}
	return result
}

func main() {
	var result int
	startTime := time.Now()
	fmt.Println("Problem 5")

	//result = smallestPositiveNumberDivisibleByNumbersLessThan(20)
	result = refinedSmallestPositiveNumberDivisibleByNumbersLessThan(20)

	fmt.Printf("Answer: %d\n", result)
	fmt.Printf("Took: %v\n", time.Now().Sub(startTime))
}
