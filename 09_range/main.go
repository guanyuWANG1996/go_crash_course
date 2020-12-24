package main

import "fmt"

func main() {
	ids := []int{1,2,3,4,5,6,7}
	sum := 0
	for i,id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
		sum += id
	}
	fmt.Println(sum)


}
