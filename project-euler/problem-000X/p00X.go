package main

import (
  "fmt"
  "time"
)

func main() {
  var result int
  startTime := time.Now()
  fmt.Println("Problem ")

  fmt.Printf("Answer: %d\n",result)

  fmt.Printf("Took: %v",time.Now().Sub(startTime))
}
