package main

import (
	"fmt"
)

func main() {
	// //1.判断字符串中汉字的数量
	// //难点是判断一个字符是汉字
	// s := "abcd我爱你"
	// //1.依次拿到字符串中的字符
	// //2.判断当前这个字符是不是汉字
	// //3.把汉字出现的次数累加
	// var count int
	// s1 := make([]string,0,100)
	// for _, v := range s {
	// 	if unicode.Is(unicode.Han, v) {
	// 		count++
	// 		s1 = append(s1, fmt.Sprintf("%c",v))
	// 	}
	// }
	// fmt.Println("汉字个数为", count)
	// fmt.Println("分别是", s1)

	// //2.How do U do 单词出现的次数
	// //1.把字符串按照空格切割得到切片
	// s1 := "How do U do"
	// //2.遍历切片存储到一个map
	// s2 := strings.Split(s1, " ")
	// m1 := make(map[string]int)
	// for _, v := range s2 {
	// 	//如果 m1中key不存在，那么把key增加，value记为1
	// 	//如果m1中key存在，那么给对应value增加1
	// 	if _, ok := m1[v]; !ok {
	// 		m1[v] = 1
	// 	} else {
	// 		m1[v] += 1
	// 	}
	// }
	// //3.累加出现的次数
	// for k, v := range m1 {
	// 	fmt.Println(k, v)
	// }

	//回文判断
	//字符串从左往右读和从右往左杜是一样的，那么就是回文。
	//上海自来水来自海上
	//山西运煤车煤运山西
	//黄山落叶松叶落山黄
	//解题思路：
	//把字符串中的字符拿出来放到[]rune中
	s := "黄山落叶松叶落山黄"
	r := make([]string, 0, len(s))
	for _, v := range s {
		r = append(r, fmt.Sprintf("%c",v))
	}
	fmt.Println(r)
	for i := 0; i < len(r)/2; i++ {
		//r[0]   r[len(r)-1]
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return //在主函数中碰到return语句，那么整个程序就会停止
		}

	}
	fmt.Println("是回文")

}
