package main

import "fmt"

func main() {
	// Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam
	// scope yang sama

	count := 0
	increment := func() {
		fmt.Println("Increment")
		count++
	}

	increment()
	increment()
	increment()
	fmt.Println(count)
}
