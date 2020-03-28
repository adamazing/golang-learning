package fibonacci


func Fibonacci(n int) int {
  if n <= 1 {
    return n
  }
  var c,b,a int = 1,1,0

  for i:=1;i < n; i++ {
    c = b + a   // Set the current value to the sum of the previous two numbers
    a = b       // n-2 = n-1
    b = c       // n-1 = current value
  }
  return c
}
