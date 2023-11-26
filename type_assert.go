package main

import "fmt"

// Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
// Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong

func random() any {
	return true
}

func main() {
	result := random()
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Unkown", value)
	}
}
