package main

import "fmt"

// Function Type Declarations
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello " + filter(name))
}

func filterName(name string) string {
	if name == "anjing" {
		return "..."
	}else{
		return name
	}
}

func main() {
	// Function tidak hanya bisa kita simpan di dalam variable sebagai value
	// Namun juga bisa kita gunakan sebagai parameter untuk function lain
	sayHelloWithFilter("Rizal", filterName)

	// Menngunakan Function As Value
	anjing := sayHelloWithFilter
	anjing("anjing", filterName)

}
