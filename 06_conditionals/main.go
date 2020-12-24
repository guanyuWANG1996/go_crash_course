package main

import "fmt"

func main() {
	x := 25
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than than %d\n", y, x)
	}

	color := "reds"

	switch color {
	case "red":
		fmt.Println("this color is red")
	case "blue":
		fmt.Println("this color is blue")
	default:
		fmt.Println("we dont have this color")
	}
}
