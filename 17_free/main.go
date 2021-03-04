package main

import "fmt"

//
//func change(slice []int) {
//	slice[0] = 200
//}
//
//func main() {
//	//slice1 := make([]int, 5, 6)
//	////slice2 := slice1[:3]
//	//slice1[0] = 100
//	//fmt.Printf("%p\n", slice1)
//	//fmt.Println(len(slice1), cap(slice1))
//	////change(slice2)
//	////fmt.Println(slice2)
//	//
//	////需要用变量承接，因为append底层可能面临扩容
//	//slice1 = append(slice1,1,2,3,4,5,6)
//	//fmt.Println(len(slice1), cap(slice1))
//	//fmt.Printf("%p", slice1)
//	////for _, v := range slice1 {
//	////	fmt.Print(v, " ")
//	////}
//	//s := [4]int{10,20,30,40}
//	////slice := []int{10, 20, 30, 40, 50}
//	////newSlice := slice[1:3]
//	//newSlice := s[:]
//	//newSlice = append(newSlice, 60)
//	//fmt.Println(s, newSlice, cap(newSlice))
//
//	//对map按照key排序
//	//scoreMap := make(map[string]int, 100)
//	//for i := 0; i < 50; i++ {
//	//	key := fmt.Sprintf("stu%02d", i)
//	//	value := rand.Intn(100)
//	//	scoreMap[key] = value
//	//}
//	//keys := make([]string, 0, 100)
//	//for k := range scoreMap {
//	//	keys = append(keys, k)
//	//}
//	//sort.Strings(keys)
//	//for _, k := range keys {
//	//	fmt.Printf("key: %v value: %v\n", k, scoreMap[k])
//	//}
//}
type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		value := stu
		m[stu.name] = &value
		fmt.Printf("%v, %#v, %v\n", stu, &stu, &stu.name)
	}
	for k, v := range m {
		//fmt.Printf("%v, %#v\n", k, v)
		fmt.Println(k, "=>", v.age)
	}
}
