package main

import "fmt"

func main() {
	a := 5
	b := &a
	// assign b to a pointer of a
	fmt.Println(a, b) // 5 0xc00001e0a8
	fmt.Printf("%T\n",b) // *int
	// use * to read val from address
	fmt.Println(*b) // 5
	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
