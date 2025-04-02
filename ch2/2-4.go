package main

import "fmt"

func main() {
	const (
		// 常量a的值为10
		a = 10

		// 常量b,c的赋值操作省略，则二者的值均为10
		b
		c
	)

	fmt.Println(a, b, c)
}
