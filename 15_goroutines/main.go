package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	f("Gin")

	go f("goroutine")

	//go f("Starting")
	go func(msg string) {
		for i := 0; i < 3; i++ {
			fmt.Println(msg)
			time.Sleep(1 * time.Second)
		}
	}("starting")

	// Keep our main Goroutine running
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
