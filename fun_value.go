package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	// Function adalah first class citizen
	// Function juga merupakan tipe data, dan bisa disimpan di dalam variable
	goodbye := getGoodBye
	fmt.Println(goodbye("Zall"))
}
