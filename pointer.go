package main

import "fmt"

// Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
// Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain,
// sebenarnya yang dikirim adalah duplikasi value nya

type Address struct{
	Region, City, Country string
}


func main() {

	address1 := Address{"Tembalang", "Semarang", "Indonesia"}
	address2 := address1

	// Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan
	// operator & diikuti dengan nama variable nya
	// address2 := &address1	// Akan Berubah

	address2.City = "Blora"

	fmt.Println(address1)	// Tidak Berubah
	fmt.Println(address2)
}
