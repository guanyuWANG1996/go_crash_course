package main

import "fmt"

func main() {
	//using var keyword
	//var name = "WGY"
	var age int32 = 25
	var isCool = true
	isCool = false

	//shorthand
	name, email := "WGY1996", "wgy@gmail.com"
	size := 1.3

	fmt.Println(name, age, isCool, size, email)
	fmt.Printf("%T\n", name)
}
