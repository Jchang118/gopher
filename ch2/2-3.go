package main

import "fmt"

var a = 10

func main() {
	fmt.Println("不使用全局变量不会编译报错，因为编译器无法确认全局变量是否会被外部程序使用")
}
