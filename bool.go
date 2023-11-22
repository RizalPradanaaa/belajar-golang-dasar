package main

import "fmt"

func main() {
	nilaiAbsen := 90
	nilaiUjian := 80

	var lulusNilaiAbsen bool = nilaiAbsen > 80
	var lulusNilaiUjian bool = nilaiUjian > 80

	var lulus bool = lulusNilaiAbsen && lulusNilaiUjian

	fmt.Println(lulus) // false
}
