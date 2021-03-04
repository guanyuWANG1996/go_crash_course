package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	// 不知道别人调用这个函数时会传进来什么类型的变量
	// 1. 方式1：断言，但需要手动猜测判断
	// 2. 方式2：借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj.Name(), obj.Kind())
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v\n", v)
}

type Person struct {
	name string
	age  int8
}

func main() {
	var a float32 = 1.234
	reflectType(a)
	reflectValue(a)
	var b int16 = 100
	reflectType(b)
	reflectValue(b)
	p := Person{
		"guanyu",
		24,
	}
	reflectType(p)
	reflectValue(p)
}
