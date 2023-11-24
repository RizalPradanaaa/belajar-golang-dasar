package main

import "fmt"

// Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data
// lainnya dalam satu kesatuan
// Struct biasanya representasi data dalam program aplikasi yang kita buat
// Data di struct disimpan dalam field
// Sederhananya struct adalah kumpulan dari field

// Template Struct
type Customer struct{
	Name, Addres string
	Age int
}

// Struct Method
func (customer Customer) sayHello(){
	fmt.Println("Hello ", customer.Name)
}

func main() {

	var rizal Customer
	rizal.Name = "Rizal"
	rizal.Addres = "Semarang"
	rizal.Age = 20

	fmt.Println(rizal)


	// Membuat Struct Dengan Program Struct Literals
	joko := Customer{
		Name: "Joko",
		Addres: "Bandung",
		Age: 30,
	}
	fmt.Println(joko)

	Budi := Customer{"Budi", "Jakarta", 25}
	Budi.sayHello()
	fmt.Println(Budi)


	// Struct Method
	Boy := Customer{"Boy", "Jakarta", 25}
	Boy.sayHello()

}
