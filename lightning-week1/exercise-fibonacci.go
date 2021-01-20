package main

import "fmt"
import "math/big"

// fibonacci is a function that returns
// a function that returns a big int.
func fibonacci() func() big.Int {
  a := big.NewInt(0)
  b := big.NewInt(1)

  return func() big.Int {
    defer func() {
      var c big.Int
      c.Add(a,b)
      a, b = b, c
    }()
    return a
  }
}

func main() {
  var f func() big.Int
	f = fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}
