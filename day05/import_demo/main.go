package main

import (
	"fmt"

	abc "github.com/a754962942/Gotest/day05/package"
)

const (
	PI = 3.1415926
	a  = 100
)

func init() {
	fmt.Println("This is main.init", PI, a)

}
func main() {
	ret := abc.Add(1, 2)
	fmt.Println(ret)
}
