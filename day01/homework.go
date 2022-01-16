package main

import "fmt"

func main() {
	//控制行数
	for i := 1; i <= 9; i++ {
		//控制列数
		for j := 1; j <= i; j++ {
			fmt.Printf("%dX%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
