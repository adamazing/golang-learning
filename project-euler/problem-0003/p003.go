package main

import (
  "fmt"
  "math"
  "math/big"
  "time"
)

// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?

func largestPrimeFactor(n int,primes []int ) int {
  for i:= len(primes)-1; i>=0;  i-- {
    if n % primes[i] == 0 {
      return primes[i]
    }
  }
  return n
}

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

func main() {
  startTime := time.Now()
  fmt.Println("Problem 3")

  sievedValues := sieve(int(math.Sqrt(600851475143)),1)

  result := largestPrimeFactor(600851475143,sievedValues)
  fmt.Printf("Answer: %d\n",result)

  fmt.Printf("\nTook: %v\n", time.Now().Sub(startTime))
}
