package main

import "fmt"

type Vertex struct {
	X int
	Y int
}



func main() {
	fmt.Println(Vertex{1, 2})
  var q = Vertex{4,5}
  fmt.Println(q.X,q.Y)
}
