package main

import (
	"fmt"
	"os"
)

/*

要求：使用方法
写一个系统可以能够查看、新增学生\删除学生
*/
//创建一个结构体名为struct,用来记录每个学生的姓名
type student struct {
	id   int
	name string
}

var allStudent map[int]*student

//newStudent创建一个学生，并存入allStudent这个map中
func (s *student) newStudent() {
	var id int
	var name string
	fmt.Print("请输入学生学号：")
	fmt.Scan(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scan(&name)
	s = &student{
		id:   id,
		name: name,
	}
	allStudent[id] = s
	fmt.Println("新增成功")

}

//查看所有student
func (s *student) showAllstudent() {
	for _, v := range allStudent {
		fmt.Printf("学号：%d\t姓名:%s\n", v.id, v.name)
	}
}

//根据学号删除指定学生
func (s *student) del() {
	var id int
	fmt.Print("请输入学生学号：")
	fmt.Scan(&id)
	if allStudent[id] != nil {
		delete(allStudent, id)
	} else {
		fmt.Println("没有该学生,删除失败")
	}

}
func main() {
	allStudent = make(map[int]*student)
	s := new(student)
	//1.查询学生
	//2.增加学生
	//3.删除学生
	for {
		fmt.Println("欢迎使用方法版学生管理系统")
		fmt.Println(`
		请选择您想要的功能
		1.查询学生
		2.新增学生
		3.删除学生
		4.退出系统
		`)
		var choice int
		fmt.Scan(&choice)
		fmt.Printf("您选择的是%d\n", choice)
		switch choice {
		case 1:
			s.showAllstudent()
		case 2:
			s.newStudent()
		case 3:
			s.del()
		case 4:
			os.Exit(1)
		}
	}
}
