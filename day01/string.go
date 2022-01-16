package main

//字符串
import (
	"fmt"
	"strings"
)

func main() {
	path := "F:\\dhw"
	fmt.Println(path)
	//多行字符串
	s2 := `
123
	345
673838
	`
	fmt.Println(s2)
	name := "abc"
	world := "dss"
	//fmt.Println(name+world)
	//拼接后返回给变量
	s1 := fmt.Sprintf("%s%s", name, world)
	fmt.Printf(s1)
	//strings.Split 分割字符串
	ret := strings.Split(s2, "3")
	fmt.Println(ret)
	//strings.contains 判断字符串是否包含
	fmt.Println(strings.Contains(s1, "a"))
	fmt.Println(strings.Contains(s1, "p"))
	//strings.HasPrefix判断字符串前缀
	fmt.Println(strings.HasPrefix(s1, "a"))
	//strings.HasSuffix判断字符串后缀
	fmt.Println(strings.HasSuffix(s1, "s"))
	ss := "abcdeb"
	//取出字符在字符串的位置
	fmt.Println(strings.Index(ss, "d"))
	//取出字符串最后出现的位置
	fmt.Println(strings.LastIndex(ss, "b"))
	//拼接字符串
	fmt.Println(strings.Join(ret, "+"))
}
