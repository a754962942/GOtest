package main

import (
	"fmt"
)

/*
有50枚金币，需要分给：Mattew、Sarah、Augustus、Heidi、Emilie、Peter、Giana、Adriano、Aaron、Elizabeth
分配规则：
1. 名字中每包含1个'e'或'E'分1枚金币
2. 名字中每包含1个'i'或'I'分2枚金币
3. 名字中每包含1个'o'或'O'分3枚金币
4. 名字中每包含1个'u'或'U'分4枚金币

写一个程序，计算每个用户分到多少金币，以及最后剩余多少

*/
var (
	coins = 50
	users = []string{
		"Mattew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println(distribution)

}
func dispatchCoin() (left int) {
	//1.依次拿到每个人的名字
	//2.拿到一个人名根据分金币的规则去分金币
	//2.1每个人分的金币应该保存到distribution中
	//2.2 还要记录下剩余的金币数
	//3.整个第2步执行完就能得到最终每个人分的金币数和剩余金币数
	left = 50

	for _, v := range users {
		distribution[v] = 0
		for _, k := range v {
			if k == 'e' || k == 'E' {
				distribution[v] += 1
				left -= 1
			} else if k == 'i' || k == 'I' {
				distribution[v] += 2
				left -= 2
			} else if k == 'o' || k == 'O' {
				distribution[v] += 3
				left -= 3
			} else if k == 'u' || k == 'U' {
				distribution[v] += 4
				left -= 4
			}
		}

	}
	return left
}
