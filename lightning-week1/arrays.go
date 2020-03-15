package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	fmt.Println(primes)

  primes = primes[0:6]
  fmt.Println(primes)

  var s []int = primes [1:4]
  fmt.Println(s)

  q := primes[0:2]
  r := primes[1:3]
  q[0]=19   // slices point to the original array

  fmt.Println(q,r)
  fmt.Println(primes)

  var f []int
	fmt.Println(f, len(f), cap(f))
	if f == nil {
		fmt.Println("nil!")
	}

  g := make([]int, 5)
  printSlice("g", g)

  gg := make([]int,0,5)
  printSlice("gg",gg)

  h := gg[:2]
  printSlice("h",h)

  i := h[2:5]
  printSlice("i",i)
}


func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
