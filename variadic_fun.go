package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	// Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs
	// Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja semacam Array
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)

	// Kadang ada kasus dimana kita menggunakan Variadic Function, namun memiliki variable berupa slice
	// Kita bisa menjadikan slice sebagai vararg parameter
	numbers := []int{10,10,10}
	fmt.Println(sumAll(numbers...))
}
