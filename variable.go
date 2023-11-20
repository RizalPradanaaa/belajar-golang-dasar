package main

import "fmt"

func main() {
	// Cara Deklarasi Pertama
	var name1 string
	name1 = "Rizal"
	fmt.Println("Coba Var Pertama = ", name1)

	// bisa tanpa menyebut tipe data, tapi harus langsung diinisiasi
	var name2 = "Nawang"
	fmt.Println("Coba Var Kedua = ", name2)

	// bisa tanpa meggunakan var, tapi harus langsung diinisiasi
	name3 := "Pradana"
	fmt.Println("Coba Tanpa Var = ",name3)

	// Deklarasi Multiple variable
	var(
		coba1 = "coba1"
		coba2 = "coba2"
	)
	fmt.Println("Coba Multiple Var = ",coba1)
	fmt.Println("Coba Multiple Var = ",coba2)


	// Const
	const cobaconst1 string = "cobaconst1"
	const cobaconst2 = "cobaconst2"

	fmt.Println("Coba Const1 = ",cobaconst1)
	fmt.Println("Coba Const1 = ",cobaconst2)

	// Deklarasi Multiple const
	const(
		cobamultiple1 = "constmulti1"
		cobamultiple2 = "constmulti2"
	)
	fmt.Println("Coba Multiple Const = ", cobamultiple1)
	fmt.Println("Coba Multiple Const = ", cobamultiple2)
}
