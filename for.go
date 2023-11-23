package main

import "fmt"

func main() {

	count := 1
	for count <= 10 {
		fmt.Println("Perulangan ke", count)
		count++
	}


	// For dengan Statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}


	names := []string{"Rizal", "Nawang", "Pradana"}
	for _, name := range names {
		fmt.Println("Nama = ", name)
	}


	// Break
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan Ke,", i)
	}

	// Continue

	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("Perulangan Ganjil Ke,", i)
	}

}
