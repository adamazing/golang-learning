package fibonacci

import "testing"

func TestFibonacci0(t *testing.T){
  got := Fibonacci(0)
  if got != 0 {
    t.Errorf("Expected 0 and got %d\n", got)
  }
}

func TestFibonacci1(t *testing.T){
  got := Fibonacci(1)
  if got != 1 {
    t.Errorf("Expected 1 and got %d\n", got)
  }
}

func TestFibonacci2(t *testing.T){
  got := Fibonacci(2)
  if got != 1 {
    t.Errorf("Expected 1 and got %d\n", got)
  }
}

func TestFibonacci3(t *testing.T){
  got := Fibonacci(3)
  if got != 2 {
    t.Errorf("Expected 2 and got %d\n", got)
  }
}

func TestFibonacci4(t *testing.T){
  got := Fibonacci(4)
  if got != 3 {
    t.Errorf("Expected 3 and got %d\n", got)
  }
}
