package main

import "fmt"

func main() {
	//Arrays
	//var fruitArr [2]string

	//Assign values
	//fruitArr[0] = "Apple"
	//fruitArr[1] = "Orange"

	//Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}
	fmt.Printf("%T\n", fruitArr)
	//Slice
	fruitSlice := []string{"Apple", "Orange", "Grape"}
	newSlice := fruitSlice[1:]
	//fmt.Printf("%T\n", fruitSlice)
	fmt.Printf("%T\n", newSlice)
	fmt.Println(newSlice[0])
}
