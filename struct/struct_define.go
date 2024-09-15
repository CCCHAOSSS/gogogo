package main

import "fmt"

// 定义结构体
type person struct {
	name string
	age  int
	sex  string
	city string
}

func main() {
	var limei person
	limei.name = "Limei"
	limei.age = 20
	limei.sex = "female"
	limei.city = "hangzhou"

	//匿名结构体
	var lili struct {
		name string
		age  int
	}
	lili.name = "Lili"
	lili.age = 20
	//fmt.Println(limei)
	//fmt.Println(lili)
	//fmt.Printf("%T\n\n", limei)
	//fmt.Printf("%p\n", limei)
	fmt.Printf("limei=%#v\n", &limei)
	fmt.Printf("lili=%#v\n", &lili)

	//指针类型结构体
}
