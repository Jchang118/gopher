package main

import "fmt"

func main() {

	var a, b int
	a = 10
	b = 20

	fmt.Println("原始值: a = ", a, ", b = ", b)

	a = a + b
	b = a - b
	a = a -b

	fmt.Println("执行后: a = ", a, ", b = ", b)
}
