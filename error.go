package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi Dengan Nol")
	}else{
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := pembagian(10, 10)
	if err == nil {
		fmt.Println("Hasil Pembagian", hasil)
	}else{
		fmt.Println("Terjadi Error", err)
	}
}
