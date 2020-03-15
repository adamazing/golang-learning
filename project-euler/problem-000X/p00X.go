package main

import (
  "fmt"
  "time"
)

func main() {
  startTime := time.Now()
  fmt.Println("Problem ")



  fmt.Printf("Took: %d",time.Now().Sub(startTime))
}
