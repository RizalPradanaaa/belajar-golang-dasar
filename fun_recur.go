package main

import "fmt"

// Factorial Dengan For Loop
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}


// Factorial Dengan Recursive Fun
func factorialRecurive(value int) int {
	if value == 1{
		return 1
	}else{
		return value * factorialRecurive(value-1)
	}
}

func main() {
	// Recursive function adalah function yang memanggil function dirinya sendir
	fmt.Println(factorialLoop(9))
	fmt.Println(factorialRecurive(9))
}
