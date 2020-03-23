package primesieve

import (
	"math"
	"math/big"
)

// Sieve function
//   n            The upper limit.
//   lowerLimit   The lower limit.
func Sieve(n, lowerLimit int) []int {
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

// NthPrime function
//   n            The upper limit.
//   lowerLimit   The lower limit.
func NthPrime(n int) int {
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
