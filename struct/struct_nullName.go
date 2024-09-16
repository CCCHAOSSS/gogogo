package main

import "fmt"

// 用匿名字段，每个只能有一个类型
type student struct {
	name string
	age  int
}

// 嵌套结构体,正常嵌套
type class struct {
	class   string
	student student
	count   int
}

// 嵌套结构体，匿名嵌套
type class2 struct {
	class string
	count int
	student
}

func main() {
	//s1 := student{
	//	"limei",
	//	20,
	//}

	//fmt.Println(s1.int) //访问时候也只能通过类型去访问

	cls1 := class{
		"class01",
		student{
			"limei",
			20,
		},
		1,
	}
	fmt.Printf("%#v\n", cls1)
	fmt.Println(cls1.student.name) //limei

	cls2 := class2{
		"classo1",
		2,
		student{
			"limei",
			20,
		},
	}
	fmt.Printf("%#v\n", cls2)
	fmt.Println(cls2.name) //limei，可以直接用匿名字段的类型去访问
}
