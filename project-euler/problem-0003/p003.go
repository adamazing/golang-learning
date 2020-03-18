package main

import (
  "fmt"
  "math"
  "time"
  "../lib/primesieve"
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

func main() {
  startTime := time.Now()
  fmt.Println("Problem 3")

  sieve := primesieve.Sieve(int(math.Sqrt(600851475143)),1)

  result := largestPrimeFactor(600851475143,sieve)
  fmt.Printf("Answer: %d\n",result)

  fmt.Printf("\nTook: %v\n", time.Now().Sub(startTime))
}
