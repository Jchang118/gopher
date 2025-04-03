package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	fmt.Println("索引位置2的元素为：", a[2])

	// 为索引位置2上的元素重新赋值
	a[2] = 0
	fmt.Println(a)
}
