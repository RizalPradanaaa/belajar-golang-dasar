package main

import "fmt"
func main() {
	// Tipe data Slice adalah potongan dari data Array
	names := [...]string{"Rizal", "Nawang", "Pradana", "Budi", "Eko", "Joko"}
	slice := names[4:6] // [Eko, Joko]
	fmt.Println(slice)


	// Append Di Slice
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daysSlice1 := days[5:]
	daysSlice1[0] = "Sabtu Baru"
	daysSlice1[1] = "Minggu Baru"
	fmt.Println(days) 		// ["Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu Baru", "Minggu Baru"]
	fmt.Println(daysSlice1)	// ["Sabtu Baru", "Minggu Baru"]


	daysSlice2 := append(daysSlice1, "Libur Baru")
	// Membuat sebuah array baru, karena append nya melebihi capacity
	daysSlice2[0] = "Upss"
	fmt.Println(days)
	fmt.Println(daysSlice1) // ["Sabtu Baru", "Minggu Baru"]
	fmt.Println(daysSlice2) // ["Upss", "Minggu Baru", "Libur Baru"]


	// Membuat Slice
	// 2 = length & 5 = Capacity
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Rizal"
	newSlice[1] = "Rizal"

	fmt.Println(newSlice)
	fmt.Println("Length newSlice", len(newSlice))
	fmt.Println("Capacity newSlice", cap(newSlice))



	// Copy Slice
	fromSlice := days[:]
	// Length dan Capacity Disesuaikan
	toSlice := make([]string, len(days), cap(days))
	// Tujuan copy, dari
	copy(toSlice, fromSlice);
	fmt.Println(fromSlice)
	fmt.Println(toSlice)



	// Perbedaan Array VS Slice
	iniArray := [...]int{1,2,3,4}
	iniSlice := []int{1,2,3,4}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
