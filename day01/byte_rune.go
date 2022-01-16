package main

import "fmt"

//byte类型和rune类型

//Go语言中为了处理非ASCII码的字符，定义了新的rune类型
func main() {

	s := "Hello 可恶"
	//len()取得byte字节的数量
	n := len(s)
	fmt.Println(n)
	//for i := 0; i < n; i++ {
	//	fmt.Println(s[i])
	//	fmt.Printf("%c \n",s[i])//%c字符
	//}

	//for _, v := range s { //从字符串中拿出具体字符
	//	fmt.Printf("%c\n", v) //%c 字符
	//}
	
	//字符串修改
	s1:="萝卜萝卜"
	//把s1转换成rune类型的slice
	s2:=[]rune(s1)
	//修改slice[0]
	s2[0]='白'
	//将修改的slice转换为string并保存在s3中
	s3:=string(s2)
	fmt.Println(s3)
}
