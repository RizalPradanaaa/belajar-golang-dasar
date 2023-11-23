package main

import "fmt"

func main() {
	// Tipe Data Map adalah tipe data kumpulan key-value,
	// kata kuncinya bersifat unik, tidak boleh sama
	person := map[string]string{
		"name":    "Rizal",
		"address": "Semarang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])


	// Membuat Map Dengan Make
	books := make(map[string]string)
	books["title"] = "Belajar Golang"
	books["author"] = "Rizal"
	books["wrong"] = "Upss"

	delete(books, "wrong")

	fmt.Println(books)
}
