package primesieve


import (
  "math"
  "math/big"
)

// Params
//   n            The upper limit.
//   lowerLimit   The lower limit.
func Sieve(n,lowerLimit int) []int {
  var prime = make([]bool, n+1)
  var primes []int
  // fill it with true values
  for i, _ := range prime {
    prime[i] = true
  }
  for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
    if big.NewInt(int64(i)).ProbablyPrime(10) {
      for j:= i*i; j<=n; j+=i {
        prime[j]=false
      }
    }
  }
  for i,v := range prime {
    if v && i>lowerLimit {
      primes = append(primes, i)
    }
  }
  return primes
}
