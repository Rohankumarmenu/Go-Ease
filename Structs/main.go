package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// alex := person{firstName: "Rahul", lastName: "Dravid"}
	// fmt.Println(alex)

	// var alex person
	// alex.firstName = "Virendra"
	// alex.lastName = "Sehwag"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	check := person{
		firstName: "Virendra",
		lastName:  "Sehwag",
		contactInfo: contactInfo{
			email:   "sehwag@example.com",
			zipCode: 12345,
		},
	}
	check.updatedName("Veeru")
	check.printed()
}

func (p person) printed() {
	fmt.Printf("%+v", p)
}

func (p person) updatedName(newFirstName string) {
	p.firstName = newFirstName
}
