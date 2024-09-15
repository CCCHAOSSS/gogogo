package main

import "fmt"

// 结构体指针
type person2 struct {
	name      string
	age       int
	city      string
	isMarried bool
}

func main() {
	var p = new(person2)  //new关键字得到person类型的指针
	fmt.Printf("%p\n", p) //输出：0xc0000204e0 只存地址，是个指针
	fmt.Println(p)        //&{ 0  false}
	//使用指针初始化变量,完全的写法：
	(*p).name = "name"
	(*p).age = 10
	(*p).isMarried = true
	(*p).city = "hangzhou"
	fmt.Printf("%#v\n", p)  //输出:  &main.person2{name:"name", age:10, city:"hangzhou", isMarried:true}
	fmt.Printf("%#v\n", *p) //输出:  main.person2{name:"name", age:10, city:"hangzhou", isMarried:true}
	fmt.Printf("%T\n", p)   //*main.person2
	fmt.Printf("%T\n", *p)  //main.person2

	//但是有语法糖，可以像直接操作结构体对象一样：
	var p2 = new(person2)
	p2.name = "name"
	p2.age = 10
	p2.city = "hangzhou"
	p2.isMarried = false
	fmt.Printf("%#v\n", *p2) //main.person2{name:"name", age:10, city:"hangzhou", isMarried:false}

	//去结构体的地址进行实例化
	fmt.Printf("\np3:")
	p3 := &person2{}
	fmt.Printf("%p\n", p3)  //0xc00008a5d0
	fmt.Printf("%T\n", p3)  //*main.person2  person2的指针类型
	fmt.Printf("%T\n", *p3) //main.person2
	p3.name = "p3"
	p3.age = 10
	//p3.city = "hangzhou"
	p3.isMarried = true
	(*p3).city = "hangzhou"
	fmt.Printf("%#v\n", *p3) //main.person2{name:"p3", age:10, city:"hangzhou", isMarried:true}
	fmt.Printf("%#v\n", p3)  //&main.person2{name:"p3", age:10, city:"hangzhou", isMarried:true}
}
