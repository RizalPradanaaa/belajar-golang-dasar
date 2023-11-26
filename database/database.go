package database

// Saat kita membuat package, kita bisa membuat sebuah function yang akan diakses ketika package
// kita diakses
// Ini sangat cocok ketika contohnya, jika package kita berisi function-function untuk berkomunikasi
// dengan database, kita membuat function inisialisasi untuk membuka koneksi ke database
// Untuk membuat function yang diakses secara otomatis ketika package diakses, kita cukup
// membuat function dengan nama ini

var connection string

func init() {
	connection = "Mysql"
}

func GetDatabase() string {
	return connection
}
