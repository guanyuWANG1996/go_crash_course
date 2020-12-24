package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["Wang"] = "WGY@qq.com"

	for name, email := range emails{
		fmt.Printf("The email for %s is %s\n", name, email)
	}
	value, ok := emails["Bob"]
	if ok{
		fmt.Println(value)
	} else{
		fmt.Println("there is no such element key")
	}
	fmt.Println(emails)
	delete(emails, "Bob")

	fmt.Println("=====================")

	for name, email := range emails{
		fmt.Printf("The email for %v is %v\n", name, email)
	}
	_, ok1 := emails["Bob"]
	if ok1{
		fmt.Println(value)
	} else{
		fmt.Println("there is no such element key")
	}

	//fmt.Println(emails)
	var m = map[string]int{
		"one": 1,
		"two": 2,
		"three":3,// Comma is necessary
	}
	fmt.Println(m["one"])
}
