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

func main() {
  var result int
  startTime := time.Now()
  fmt.Println("Problem 5")

  fmt.Printf("Answer: %d\n",result)

  fmt.Printf("Took: %v",time.Now().Sub(startTime))
}
