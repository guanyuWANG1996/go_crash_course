package main

import (
	"fmt"
	"github.com/guanyuWANG/go_crash_course/03_packages/strutil"
)

func greeting(name string) string {
	return "Hello "+ name
}

func getSum(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(greeting("Guanyu"))
	fmt.Println("the sum is",getSum(6, 2))
	fmt.Println(strutil.Reverse("nihao"))
}
