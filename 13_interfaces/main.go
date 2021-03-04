package main

import "fmt"

//// Define interface
//type Shape interface {
//	area() float64
//}
//
//type Circle struct {
//	x, y, radius float64
//}
//
//type Rectangle struct {
//	width, height float64
//}
//
//func (c Circle) area() float64 {
//	return math.Pi * c.radius * c.radius
//}
//
//func (r Rectangle) area() float64 {
//	return r.width * r.height
//}
//
////func getArea(s Shape) float64 {
////	return s.area()
////}
//
//func main() {
//	var circle Shape = Circle{
//		x:      0,
//		y:      0,
//		radius: 5,
//	}
//	var rect Shape = Rectangle{
//		width:  10,
//		height: 10,
//	}
//	fmt.Printf("the area of circle is %f\n", circle.area())
//	fmt.Printf("the area of rectangle is %f\n", rect.area())
//}
type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = &Student{}
	peo1 := &Student{}
	think := "bitch"
	think1 := "sb"
	fmt.Println(peo.Speak(think))
	fmt.Println(peo1.Speak(think1))
}
