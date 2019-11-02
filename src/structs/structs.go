package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	p := person{firstName: "I", lastName: "P"}
	fmt.Println(p.firstName)
	fmt.Println(p.lastName)
}
