package main

import "fmt"

//结构体模拟实现其他语言中的继承
type animal struct {
	name string
}
type dog struct {
	feet uint
	animal
}

func (d dog) wang() {
	fmt.Printf("%s在叫：汪汪汪\n", d.name)
}
func (a animal) move() {
	fmt.Printf("%s move\n", a.name)
}
func main() {
	d := dog{
		feet: 4,
		animal: animal{
			name: "jdj",
		},
	}
	fmt.Println(d)
	d.wang()
	d.move()
}
