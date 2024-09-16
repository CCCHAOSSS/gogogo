package main

import "fmt"

type Student struct {
	Name, Sex string
	Id, Age   int
}

// Student的构造函数
func newStudent(name, sex string, id, age int) *Student {
	return &Student{
		Name: name,
		Sex:  sex,
		Id:   id,
		Age:  age,
	}
}

// 学生管理
type StudentManager struct {
	allStudents []*Student //切片类型
}

// 学生管理的构造函数，先初始化内存
func NewStudentManager() *StudentManager {
	return &StudentManager{
		allStudents: make([]*Student, 0, 100),
	}
}

// 1.添加学生：往学生管理中添加
func (s *StudentManager) addStudent(newStu *Student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 2.编辑学生
func (s *StudentManager) updateStudent(stu *Student) {
	//for _, v := range s.allStudents {
	//	if v.Id == stu.Id {
	//		v.Name = stu.Name
	//		v.Id = stu.Id
	//		v.Age = stu.Age
	//		v.Sex = stu.Sex
	//	}
	//	return
	//}
	for i, v := range s.allStudents {
		if v.Id == stu.Id {
			s.allStudents[i] = stu
			return
		}
	}
	//走完了但没找到
	fmt.Printf("输入有误，未找到学号为%s的学生\n", stu.Id)
}

// 3.展示所有学生
func (s *StudentManager) showAllStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d\t姓名：%s\t性别：%s\t年龄：%d\n", v.Id, v.Name, v.Sex, v.Age)
	}
}
