package main

import "fmt"

func main() {
	var a = 65536
	fmt.Println("值类型a的内容：", a)

	var b = 65536
	var p = &b
	fmt.Println("指针类型p的内容：", p)
}
