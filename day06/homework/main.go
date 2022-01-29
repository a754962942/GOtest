package main

import (
	"fmt"
	"reflect"
)

//ini文件解析器

//MySqlConfig MySQL配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini":password`
}

func loadIni(fileName string, v interface{}) (err error) {
	//0. 参数校验
	//0.1 传进来的必须是指针类型
	t := reflect.TypeOf(v)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err := fmt.Errorf("传入参数错误")
		return err
	}
	//0.2传进来的指针必须是结构体指针
	// //1.读文件得到字节类型的数据

	//2.一行一行的读文件
	//2.1如果是注释就跳过
	//2.2如果是[开头的就表示节
	//2.3如果不是[开头就是=分割的键值对

	return
}
func main() {
	var mc MysqlConfig
	// loadIni(&mc)
	// fmt.Println(mc)
	err := loadIni("D:\\GO\\src\\github.com\\a754962942\\Gotest\\day06\\homework\\config.ini", &mc)
	if err != nil {
		fmt.Printf("load Ini failed,err:\n", err)
	}
}
