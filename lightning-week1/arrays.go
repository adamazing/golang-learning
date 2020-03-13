package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

  var s []int = primes [1:4]
  fmt.Println(s)

  q := primes[0:2]
  r := primes[1:3]
  q[0]=19   // slices point to the original array

  fmt.Println(q,r)
  fmt.Println(primes)
}
