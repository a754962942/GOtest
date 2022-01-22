package main

import (
	"fmt"
)

func assign(a interface{}) {
	fmt.Printf("type:%T \n", a)

	switch t := a.(type) {
	case int:
		fmt.Println("是一个Int", t)
	case string:
		fmt.Println("是一个String", t)
	case bool:
		fmt.Println("是一个Bool", t)
	}
}
func main() {
	assign(1)
}
