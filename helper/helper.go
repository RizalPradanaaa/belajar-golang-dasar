package helper

var version = 1.0 // Tidak Bisa Diakses dari luar
var Application = "Golang"

// Tidak Bisa Diakses dari luar
func sayGodbye(name string) string {
	return "Goodbye " + name
}

func SayHello(name string) string {
	return "Hello " + name
}
