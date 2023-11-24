package main

import "fmt"

type Blacklist func(string) bool

func userRegister(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are Blocked")
	}else{
		fmt.Println("Register Success")
	}
}

func main() {
	// Lebih mudah membuat function secara langsung di variable atau
	// parameter tanpa harus membuat function terlebih dahulu
	// Hal tersebut dinamakan anonymous function, atau function tanpa nama

	// Anonymous Function Ssbagai Value
	blocked := func(name string) bool {
		return name == "anjing"
	}
	userRegister("Rizal", blocked)

	// Anonymous Function Sebagai Parameter
	userRegister("anjing", func(s string) bool {
		return s == "anjing"
	})

}
