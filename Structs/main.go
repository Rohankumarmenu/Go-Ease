package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// alex := person{firstName: "Rahul", lastName: "Dravid"}
	// fmt.Println(alex)

	var alex person
	alex.firstName = "Virendra"
	alex.lastName = "Sehwag"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)

}
