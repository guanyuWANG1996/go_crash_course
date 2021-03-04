package main

import (
	"fmt"
	"strings"
)

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func makeSuffixFunc(suffix string) func(word string) string {
	return func(word string) string {
		if !strings.HasSuffix(word, suffix) {
			return word + suffix
		}
		return word
	}
}

func main() {
	//sum := adder()
	func() {
		fmt.Println("test closures")
	}()
	//for i := 0; i < 10; i++ {
	//	fmt.Println(sum(i))
	//}
	test := makeSuffixFunc("啦啦啦")
	fmt.Println(test("冠宇啦啦啦"))

	defer func() {
		err := recover()
		if err == nil {
			fmt.Println("no panic")
		} else {
			fmt.Println("panic")
		}
	}()
	panic("eeeeeee")
}
