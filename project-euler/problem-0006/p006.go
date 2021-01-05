package main

// The sum of the squares of the first ten natural numbers is,
//
// 1^2+2^2+...+10^2=385
// The square of the sum of the first ten natural numbers is,
//
// (1+2+...+10)^2=55^2=3025
// Hence the difference between the sum of the squares of the first ten natural numbers
// and the square of the sum is 3025−385=2640.
//
// Find the difference between the sum of the squares of the first one hundred natural
// numbers and the square of the sum.
import (
	"fmt"
	"math"
	"time"
)

func valsToLimit(limit int) []int {
	var valsUnderLimit = make([]int, limit)
	for i := range valsUnderLimit {
		valsUnderLimit[i] = i + 1
	}
	return valsUnderLimit
}

func sumSquares(vals []int) int {
	result := 0
	for _, v := range vals {
		result += int(math.Pow(float64(v), 2))
	}
	return result
}

func squareSums(vals []int) int {
	result := 0
	for _, v := range vals {
		result += v
	}
	return int(math.Pow(float64(result), 2))
}

func sumSquareDiff(n int) int {
	vals := valsToLimit(n)
	return squareSums(vals) - sumSquares(vals)
}

func main() {
	var result int
	startTime := time.Now()
	fmt.Println("Problem 6")

	result = sumSquareDiff(100)

	fmt.Printf("Answer: %d\n", result)
	fmt.Printf("Took: %v\n", time.Now().Sub(startTime))
}
