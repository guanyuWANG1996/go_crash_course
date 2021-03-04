package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	Address
	gender string
	age    int
}

type Address struct {
	city     string
	province string
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "Male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{
		firstName: "Guanyu",
		lastName:  "Wang",
		Address:   Address{city: "Beijing", province: "China"},
		gender:    "Female",
		age:       25,
	}
	var person2 = &Person{}
	fmt.Printf("%T\n", person2)
	fmt.Printf("%#v\n", person2)

	var user struct {
		name    string
		married bool
	}
	user.name = "WGY"
	user.married = false
	fmt.Printf("%#v\n", user)

	fmt.Printf("%#v\n", person1)
	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())
}
