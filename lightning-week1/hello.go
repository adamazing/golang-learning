package main

import (
  "fmt"
  "math/rand"
  "math"
  "time"
)

func add(x, y int) int {
  return x + y;
}

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
  fmt.Println("Hello, cruel world.👋")
  fmt.Printf("My favourite number is %d\n", rand.Intn(5876))
  fmt.Println(math.Pi)
  fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))

  var x = rand.Intn(5876)
  fmt.Printf("Doubling %d = %d\n", x, add(x,x))

  var m,n = rand.Intn(1838), rand.Intn(1893)
  fmt.Printf("Adding %d to %d equals %d \n", m, n, add(m,n))



  fmt.Println("Goodbye, cruel world.👋")
}
