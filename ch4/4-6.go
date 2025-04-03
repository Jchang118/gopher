package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	fmt.Printf("切片新建状态，长度：%d，容量：%d\n", len(s), cap(s))

	_ = append(s, 5, 6, 7, 8, 9)
	fmt.Printf("切片追加元素，忽略返回值，长度：%d，容量：%d\n", len(s), cap(s))

	s1 := append(s, 10)
	fmt.Printf("切片追加元素，赋值给新变量后，原始变量s，长度：%d，容量：%d\n", len(s), cap(s))
	fmt.Printf("切片追加元素，赋值给新变量后，新变量s1，长度：%d，容量：%d\n", len(s1), cap(s1))
}
