package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a float32

	// 未显式声明变量b的数据类型
	b := 1.4

	fmt.Println("显式声明，a的类型是：", reflect.TypeOf(a))

	// 类型推导的结果与操作系统平台位数相关
	fmt.Println("类型推导，b的类型是：", reflect.TypeOf(b))
}
