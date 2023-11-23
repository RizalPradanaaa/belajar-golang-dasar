package main

import "fmt"

// Function Dengan Paramater
func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// Function Dengan Return Value
func getHello(name string) string{
	return "Hello " + name
}

// Function Dengan Multiple Return Value
func getFullName() (string, string){
	return "Rizal", "Pradana"
}

// Function Named Return Value
func getCompleteName() (satu, dua, tiga string)  {
	satu = "Rizal"
	dua = "Nawang"
	tiga = "Pradana"

	return satu, dua, tiga
}

func main() {
	// Function Biasa
	sayHello("Rizal", "Pradana")

	// Function Return Value
	name := getHello("Rizal")
	fmt.Println(name)

	// Function Return Multiple Value
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	// Function Named Return Value
	satu, dua, tiga := getCompleteName()
	fmt.Println(satu)
	fmt.Println(dua)
	fmt.Println(tiga)
}
