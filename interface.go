package main

import "fmt"

// Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
// Sebuah interface berisikan definisi-definisi method
// Biasanya interface digunakan sebagai kontrak

type Hashname interface{
	GetName() string
}

func sayHello(hasname Hashname)  {
	fmt.Println("Hello ", hasname.GetName())
}


// Implementasi Interface
type Person struct{
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct{
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main(){
	personn := Person{Name:"Rizal"}
	sayHello(personn)
	animall := Animal{Name:"Kucing"}
	sayHello(animall)

}
