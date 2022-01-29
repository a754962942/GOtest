package main

import (
	"fmt"
	"strconv"
)

//strconv
func main() {
	str := "10000"
	//strconv.PaeseInt(转换的字符，进制，类型)
	//返回值为int64
	ret1, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		fmt.Println("parseint failed ,err:", err)
		return
	}
	fmt.Printf("%#v,%T\n", ret1, ret1) //10000,int64
	//strconv.Atoi( )返回值类型为Int
	retint, err1 := strconv.Atoi(str)
	if err != nil {
		fmt.Println("AtoI failed,err:", err1)
		return
	}
	fmt.Println(retint)
	//从字符串中解析布尔值
	boolstr := "TRUE"
	boolvalue, err := strconv.ParseBool(boolstr)
	if err != nil {
		fmt.Println("parse Bool failed,err:", err)
		return
	}
	fmt.Println(boolvalue)
	floatstr := "3.1415921"
	fstr, err := strconv.ParseFloat(floatstr, 32)
	if err != nil {
		fmt.Println("parseFloat failed,err:", err)
		return
	}
	fmt.Printf("%#v,%T\n", fstr, fstr)
	//把数字转换为字符串类型
	// i := 97
	// ret := fmt.Sprintf("%d", i) //“97”
	// fmt.Printf("%#v\n", ret)

}
