package main

import (
  "fmt"
  "time"
)

func reverseInt(number int) int {
  var remainder, reverse int = 0,0
  for{
    remainder = number%10
    reverse = reverse*10  + remainder
    number /= 10
    if number == 0 {
      break
    }
  }
  return reverse
}

func isPalindromicNumber(n int) bool{
  if n == reverseInt(n) { return true }
  return false
}

func largestPalindromicNumber (limit int) int{
  halfLimit := limit/2
  testPal := 0
  highestPal := 0
  for i:= limit; i>halfLimit; i-- {
    for j:= i; j>halfLimit; j-- {
      testPal = i*j
      if isPalindromicNumber(testPal) && testPal > highestPal {
        highestPal = testPal
      }
    }
  }
  return highestPal
}

func main() {
  startTime := time.Now()
  fmt.Println("Problem 4")

  fmt.Println(largestPalindromicNumber(999))

  fmt.Printf("Took: %v\n",time.Now().Sub(startTime))
}
