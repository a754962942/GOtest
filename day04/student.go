package main

import (
	"fmt"
	"os"
)

/*

	函数版学生管理系统
	写一个系统可以能够查看、新增学生\删除学生
*/
type student struct {
	id   int
	name string
}

var (
	allStudent map[int]*student //变量声明

)

func newStudent(id int, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}
func find() {
	//查看学生
	for k, v := range allStudent {
		fmt.Printf("学号：%d\t姓名：%s\n", k, v.name)
	}
}
func add() {
	//向allStudent中添加一个新学生
	//1.创建一个新学生
	var (
		id   int
		name string
	)
	fmt.Print("请输入学生的学号")
	fmt.Scan(&id)
	fmt.Print("请输入学生的姓名")
	fmt.Scan(&name)
	//创建一个学生(调用构造函数)
	newstu := newStudent(id, name)
	fmt.Println("创建成功")
	//把new的学生追加到allStudent中
	allStudent[id] = newstu

}
func del() {
	//显示提示信息
	var deleteId int
	fmt.Print("请输入您想删除的学生学号")
	fmt.Scan(&deleteId)
	delete(allStudent, deleteId)
	fmt.Println("删除成功")
}

func main() {
	allStudent = make(map[int]*student, 48) //初始化(开辟内存空间)
	for {

		//1.打印菜单
		fmt.Println("欢迎使用学生管理系统")
		fmt.Println(`
		1.查看所有学生
		2.新增学生
		3.删除学生
		4.退出系统
	`)
		fmt.Print("请输入您想要选择的功能：")
		//2.等待用户选择
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("您选择了:%d\n", choice)
		//3.执行对应的函数
		switch choice {
		case 1:
			find()
		case 2:
			add()
		case 3:
			del()
		case 4:
			os.Exit(1)
		}
	}
}
