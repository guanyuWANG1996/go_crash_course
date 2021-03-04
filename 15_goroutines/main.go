package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f(from string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() { //开启主goroutine去执行main函数
	wg.Add(2)
	go f("goroutine")
	//go f("Starting")
	go func(msg string) {
		for i := 0; i < 3; i++ {
			fmt.Println(msg)
		}
		wg.Done()
	}("starting")
	fmt.Println("Main function")
	wg.Wait()
}
