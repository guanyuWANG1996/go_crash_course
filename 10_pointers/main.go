package main

import "fmt"

//func main() {
//	a := 5
//	b := &a
//	// assign b to a pointer of a
//	fmt.Println(a, &a) // 5 0xc00001e0a8
//	fmt.Printf("%T\n",b) // *int
//	fmt.Printf("%v %p\n", b, &b)
//	// use * to read val from address
//	fmt.Println(*b) // 5
//	// Change val with pointer
//	*b = 10
//	fmt.Println(a)
//}

func main() {
	var a *int //声明空指针
	a = new(int)
	*a = 100
	fmt.Println(*a)

	b := make(map[string]int, 10)
	b["沙河娜扎"] = 100
	fmt.Println(b)
}
