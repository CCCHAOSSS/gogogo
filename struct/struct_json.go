package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	age  int
	Id   int
}
type Class struct {
	Title   string
	Student []Student
}

func main() {
	c := Class{
		"101",
		make([]Student, 10),
	}

	for i := 0; i < len(c.Student); i++ {
		stu := Student{
			Name: fmt.Sprintf("student%d", i),
			age:  20,
			Id:   i,
		}
		c.Student[i] = stu
	}

	//json序列化
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("json:%s\n", data)

	//反序列化
	str := "{\"Title\":\"101\",\"Student\":[{\"Name\":\"student 0\",\"age\":20,\"Id\":0},{\"Name\":\"student 1\",\"age\":20,\"Id\":1}]}"
	c1 := &Class{}
	//解析出来，放到c1里，要传入指针才能修改c1
	err = json.Unmarshal([]byte(str), c1)

	//或者这样写，先声明，然后传入指针
	var c2 Class
	//Class是一个结构体，这是值类型，var时候会自动分配了内存
	//引用类型：指针、切片、map
	err = json.Unmarshal([]byte(str), &c2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%#v\n", *c1)

	//输出想要的值
	fmt.Println(c2.Student[0].Name)
}
