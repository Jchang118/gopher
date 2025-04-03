package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	fmt.Printf("切片s底层数组地址：%p\n", s)

	s1 := s[1:]
	fmt.Printf("切片s1底层数组地址：%p\n", s1)
}
