package main

import (
  "fmt"
  "math"
  "math/big"
  "time"
)

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
		if big.NewInt(int64(i)).ProbablyPrime(20) {
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

func sum(primes ...int) (result int) {
  for _, prime := range primes {
    result += prime
  }
  return
}

func sumPrimesBelow(n int) int {
  return sum(sieve(n,1)...)
}

func main() {
  var result int
  startTime := time.Now()
  fmt.Println("Problem 10")

  result = sumPrimesBelow(2_000_000)
  fmt.Printf("Answer: %d\n",result)

  fmt.Printf("Took: %v\n",time.Now().Sub(startTime))
}
