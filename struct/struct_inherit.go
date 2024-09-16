package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

type Cat struct {
	Feet    int
	*Animal //匿名嵌套 结构体指针实现
}

func (c *Cat) miao() {
	fmt.Printf("%s会喵喵~\n", c.name)
}

func main() {
	limei := &Cat{
		Feet:   4,
		Animal: &Animal{"limei"},
	}
	(*limei).move()
	limei.miao()
	fmt.Printf("%T\n", limei)

	lili := Cat{
		Feet:   4,
		Animal: &Animal{"lili"},
	}
	lili.move()
	lili.miao()
	fmt.Printf("%T\n", lili)
}
