package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	// 调用change()函数前，打印数组内容
	fmt.Println("调用函数前：", a)

	// 调用change()函数，尝试修改数组元素的值
	change(a)

	// 调用change()函数后，打印数组内容
	fmt.Println("调用函数后：", a)
}

func change(a [5]int) {
	a[2] = 0

	// 在函数中打印数组参数的内容
	fmt.Println("函数执行中：", a)
}

