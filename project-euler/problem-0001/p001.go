package main

import (
  "fmt"
  "time"
)

// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.

func main() {
  startTime := time.Now()
  fmt.Println("Problem 1")
  total := 0
  for i:=1; i<1000; i++ {
    if i%3 == 0 {
      total+=i
    }else if i%5 == 0 {
      total+=i
    }
  }
  fmt.Printf("Total is %d\n", total)

  fmt.Printf("Took: %v",time.Now().Sub(startTime))
}
