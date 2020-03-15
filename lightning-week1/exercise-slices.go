package main

import (
  // "golang.org/x/tour/pic"
  "math/rand"
  "fmt"
)

func Pic(dx, dy int) [][]uint8 {
  pic := make([][]uint8,dy)
  for i, _ := range pic {
    pic[i] = make([]uint8,dx)
	}
  for j, v := range pic {
    for k, _ := range v {
      //pic[j][k] = uint8(j)^uint8(k)
      //pic[j][k] = uint8(j)*uint8(k)
      pic[j][k] = uint8(j^k)-uint8(k^j) / 2
    }
  }
  return pic
}

func main() {
	//pic.Show(Pic)
  picArray := Pic(12,12)
  for _,v := range picArray {
    fmt.Println(v)
  }

}
