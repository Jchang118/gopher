package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	fmt.Printf("切片新建状态，底层数组地址：%p\n", s)

	s = append(s, 5, 6, 7, 8, 9)
	fmt.Printf("切片追加元素，未超过当前容量，底层数组地址：%p\n", s)

	s = append(s, 10, 111)
	fmt.Printf("切片追加元素，已超过当前容量，底层数组地址：%p\n", s)
}
