package main

import "fmt"

type personM struct {
	name, city string
	age        int
}

func newPersonM(age int, name, city string) *personM {
	return &personM{
		age:  age,
		name: name,
		city: city,
	}
}

func newPersonAll(age int, name, city string) personM {
	return personM{
		age:  age,
		name: name,
		city: city,
	}
}
func (p personM) Dream() {
	//p.name = "ByteDancer"
	fmt.Printf("%s在%s字节跳动工作！\n", p.name, p.city)
}

func (p personM) Dream2() {
	p.city = "shanghai"
	fmt.Printf("%s在%s字节跳动工作！\n", p.name, p.city)
}

func (p personM) years() (count int) {
	x := 60 - p.age
	return x
}

func (p *personM) setDreams(age int, city string) {
	p.age = age
	p.city = city
}

func (p personM) setDreams2(age int, city string) {
	p.age = age
	p.city = city
}

func main() {
	//p1是一个指针类型
	p1 := newPersonM(10, "limei", "hangzhou")
	p1.name = "zhazha"

	//p2是一个person的struct
	p2 := newPersonAll(20, "lili", "hangzhou")

	fmt.Printf("p1为%T\n", p1)
	fmt.Printf("这里使用构造函数(是个函数)构造p1:\n%#v\n", *p1)
	fmt.Println("下一句是使用方法输出")
	p1.Dream()

	fmt.Printf("\n")

	fmt.Printf("p2为%T\n", p2)
	fmt.Printf("使用构造函数构造p2:\n%#v\n", p2)
	fmt.Println("下一句是使用方法输出")
	p2.Dream2()

	years := p2.years()
	fmt.Printf("%s还有%d年退休\n", p2.name, years)

	//
	p3 := newPersonM(100, "sily", "guangzhou")
	p4 := newPersonM(100, "babo", "guangzhou")

	//方法接收者为指针的，能被方法修改值：
	p3.setDreams(99, "beijing")
	fmt.Println(*p3) //{sily beijing 99}

	//方法为struct的，无法被修改
	//因为func里面传参传的是值拷贝
	p4.setDreams2(99, "beijing")
	fmt.Println(*p4) //{babo guangzhou 100}

}
