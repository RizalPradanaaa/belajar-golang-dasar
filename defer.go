package main

import (
	"fmt"
)

// Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function
// selesai di eksekusi
// Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
func logging() {
	fmt.Println("Selesai")
}
func runApplication() {
	// fun logging akan selalu dijalankan saat fun runapp selesai
	defer logging()
	fmt.Println("App Running")
}


// Panic function adalah function yang bisa kita gunakan untuk menghentikan program
// Panic function biasanya dipanggil ketika terjadi panic pada saat program kita berjalan
// Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

// Recover adalah function yang bisa kita gunakan untuk menangkap data panic
// Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan
func endApp()  {
	fmt.Println("App End")
	// fun Recover Melanjutkan program yang panic
	message := recover()
	fmt.Println("Terjadi Error", message)
}

func runApp(error bool)  {
	// fun panic, akan menghentikan program, tapu defer tetap dijalankan
	defer endApp()
	if error {
		panic("ERROR")
	}
}
func main() {


	//runApplication()
	runApp(true)

}
