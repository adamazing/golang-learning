package main

import (
  "fmt"
  "math"
  "time"
)

func getFirstPythagoreanTripletThatSumsToN(sum int64) int64 {
  var i, j int64
  for i=1;i<sum;i++ {
    for j=i;j<sum;j++ {
      testZ := math.Pow(float64(i),2) + math.Pow(float64(j),2)
      testZRoot := math.Sqrt(testZ)
      if testZRoot == math.Round(testZRoot) {
        if float64(i) + float64(j) + testZRoot == float64(sum) {
          return (i * j * int64(testZRoot))
        }
      }
    }
  }
  return 0
}

func main() {
  var result int64
  startTime := time.Now()

  result = getFirstPythagoreanTripletThatSumsToN(1000)
  fmt.Println("Problem 9")

  fmt.Printf("Answer: %d\n",result)

  fmt.Printf("Took: %v\n",time.Now().Sub(startTime))
}
