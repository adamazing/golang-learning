package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X float64
	Y float64
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func (v1 *vertex) euclideanDistance(v2 vertex) float64 {
	return hypot(math.Abs(v2.X - v1.X), math.Abs(v2.Y - v1.Y))
}

func (v1 *vertex) manhattanDistance(v2 vertex) float64 {
	return math.Abs(v2.X - v1.X) + math.Abs(v2.Y - v1.Y)
}

func main() {
  var q = vertex{0,0}
  var r = vertex{3,4}
  fmt.Println(q.X,q.Y)
  fmt.Println(r.X,r.Y)

	fmt.Println(q.euclideanDistance(r))
	fmt.Println(r.manhattanDistance(q))
}
