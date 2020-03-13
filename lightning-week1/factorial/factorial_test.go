package factorial

import "testing"

func TestFactorial0(t *testing.T){
  got := Factorial(0)
  if got != 1 {
    t.Errorf("Expected 1 and got %d\n", got)
  }
}

func TestFactorial1(t *testing.T){
  got := Factorial(1)
  if got != 1 {
    t.Errorf("Expected 1 and got %d\n", got)
  }
}

func TestFactorial2(t *testing.T){
  got := Factorial(2)
  if got != 2 {
    t.Errorf("Expected 1 and got %d\n", got)
  }
}

func TestFactorial3(t *testing.T){
  got := Factorial(3)
  if got != 6 {
    t.Errorf("Expected 1 and got %d\n", got)
  }
}
