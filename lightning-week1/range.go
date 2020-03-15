package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}


func main() {
  fmt.Println(pow)
	for i, v := range pow {
    pow[i] = 2*v
	}
  fmt.Println(pow)
}
