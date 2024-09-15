package main

import "fmt"

type person3 struct {
	name, city string
	age, sex   int
}

func main() {
	//键值对初始化，可以有不赋值的
	p1 := person3{
		name: "limei",
		city: "hangzhou",
		age:  22,
		sex:  1,
	}
	fmt.Println(p1)

	//使用值的列表进行初始，必须初始化所有字段，只会按序赋值
	p2 := person3{
		"lili",
		"hangzhou",
		18,
		1,
	}
	fmt.Println(p2)

	//结构体构造函数
	//结构体是值类型，复制时候会完整复制
	//在实现构造函数时候，返回指针类型会消耗小一点
	p3 := newPerson(10, 1, "hangzhou", "lili")
	fmt.Println(*p3)
}

func newPerson(age, sex int, city, name string) *person3 {
	return &person3{
		name: name,
		city: city,
		age:  age,
		sex:  sex,
	}
}
