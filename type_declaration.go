package main

import "fmt"

func main() {
	// Membuat Tipe Declaration atau inisial
	// Sekarang NoKtp adalah tipe deklarasi string
	type noKtp string
	var ktp noKtp = "12344"
	// var ktp string = "12344"
	fmt.Println(ktp)

}
