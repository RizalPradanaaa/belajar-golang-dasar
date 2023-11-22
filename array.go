package main

import "fmt"

func main() {
	// Membuat array, dengan capacity 3
	var array1 [3]string
	array1[0] = "Rizal"
	array1[1] = "Nawang"
	array1[2] = "Pradana"

	fmt.Println(array1)
	fmt.Println(array1[0])
	fmt.Println(array1[1])
	fmt.Println(array1[2])

	// Membuat array langsung angka 3 bisa diganti ...
	// Jika ... ukuran array akan sesuai dengan jumlah data didalamnya
	// var values = [...]int{}
	var values = [3]int{
		1,
		2,
		3}
	fmt.Println(values)
	fmt.Println(len(values));
	values[0] = 100
	fmt.Println(values)
}
