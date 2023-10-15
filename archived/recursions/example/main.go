package main

import (
	"fmt"
	"github.com/nhaancs/go-algorithms/archived/recursions"
)

func main() {
	// fmt.Println(recursions.Factorial(0))
	// fmt.Println(recursions.Factorial(1))
	// fmt.Println(recursions.Factorial(2))
	// fmt.Println(recursions.Factorial(3))
	// fmt.Println(recursions.Factorial(4))
	// fmt.Println(recursions.Factorial(5))

	fmt.Println(
		recursions.FibonacciNumber(0),
		recursions.FibonacciNumber(1),
		recursions.FibonacciNumber(2),
		recursions.FibonacciNumber(3),
		recursions.FibonacciNumber(4),
		recursions.FibonacciNumber(5),
		recursions.FibonacciNumber(6),
		recursions.FibonacciNumber(7),
		recursions.FibonacciNumber(8),
		recursions.FibonacciNumber(9),
		recursions.FibonacciNumber(10),
		recursions.FibonacciNumber(11),
		recursions.FibonacciNumber(40),
	)
}
