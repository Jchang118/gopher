package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s string
	s = "Go语言字符串"

	fmt.Printf("s的内容为：%v\n", s)
	fmt.Printf("s的数据类型为：%s\n", reflect.TypeOf(s))
}
