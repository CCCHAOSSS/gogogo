package main

import (
	"fmt"
	"os"
)

func showMenu() {
	fmt.Println("欢迎来到学生管理系统")
	fmt.Println("请输入序号来使用相关功能")
	fmt.Println("1: 添加新的学生")
	fmt.Println("2: 根据学号修改学生信息")
	fmt.Println("3: 展示所有学生信息")
	fmt.Println("4: 根据学号删除学生信息")
	fmt.Println("5: 退出系统")
}

func inputInfo() *Student {
	var (
		name string
		age  int
		id   int
		sex  string
	)
	fmt.Println("按要求输入信息：")
	fmt.Println("请输入学号：")
	fmt.Scanf("%d", &id)
	fmt.Println("请输入姓名：")
	fmt.Scanf("%s", &name)
	fmt.Println("亲输入性别：")
	fmt.Scanf("%d", &sex)
	fmt.Println("请输入年龄：")
	fmt.Scanf("%d", &age)
	stu := newStudent(name, sex, id, age)
	return stu
}
func main() {
	sm := NewStudentManager()
	for {
		//1,展示菜单
		showMenu()
		//2. 等待用户执行
		fmt.Println("请输入：")
		var input int
		fmt.Scanf("%d", &input)
		//3.执行
		switch input {
		case 1:
			stu := inputInfo()
			sm.addStudent(stu)
		case 2:
			stu := inputInfo()
			sm.updateStudent(stu)
		case 3:
			sm.showAllStudent()
		case 4:
			fmt.Println("本功能还未完成")
		case 5:
			os.Exit(0)
		}
	}
}
