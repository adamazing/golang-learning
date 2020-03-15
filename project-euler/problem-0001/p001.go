package main

import (
  "fmt"
  "time"
)

func main() {
  startTime := time.Now()
  fmt.Println("Problem ")


  fmt.Println(time.Now().Sub(startTime))
}
