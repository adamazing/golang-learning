package main

import (
  "fmt"
  "math/rand"
  "math"
  "time"
)

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  fmt.Println("Hello, cruel world.👋")
  fmt.Printf("My favourite number is %d\n", rand.Intn(5876))
  fmt.Println(math.Pi)
  fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
  fmt.Println("Goodbye, cruel world.👋")
}
