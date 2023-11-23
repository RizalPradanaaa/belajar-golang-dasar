package main

import "fmt"

func main() {

	person := "wkwkww"

	switch person {
	case "Rizal":
		fmt.Println("Hello Rizal")
	case "Eko" :
		fmt.Println("Hello Eko")
	default :
		fmt.Println("Hello Siapa Namanya?")
	}


	// Switch dengan Short Statement
	switch length:= len(person); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang!")
	case false:
		fmt.Println("Nama Sudah Benar!")
	}
}
