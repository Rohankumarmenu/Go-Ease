package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Rahul", lastName: "Dravid"}
	fmt.Println(alex)
}
