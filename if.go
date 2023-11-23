package main

import "fmt"

func main() {
	person := "Rizal"

	if person == "Rizal" {
		fmt.Println("Hello Rizal")
	}else if person == "eko"{
		fmt.Println("Hello Eko")
	}else{
		fmt.Println("Hello Siapa Namanya?")
	}


	// If Short Statement
	if length := len(person); length > 5 {
		fmt.Println("Nama Terlalu Panjang!")
	}else{
		fmt.Println("Nama Sudah Benar")
	}
}
