package main

import (
	"encoding/json"
	"fmt"
)

//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string `json:"title"`
	Students []Student
}

func main() {
	c := Class{
		Title:    "101",
		Students: make([]Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//fmt.Printf("%#v\n", c)
	// json 序列化
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed", err)
		return
	}
	fmt.Printf("%s\n", data)
	//fmt.Printf("%T\n", data)
	// json 反序列化
	jsonStr := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"}]}`
	var c2 Class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json unmarshal failed", err)
		return
	}
	fmt.Printf("%#v", c2)
	fmt.Println("%T", c2)
}
