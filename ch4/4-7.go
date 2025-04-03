package main

import "fmt"

func main() {
	array := [5]int{10, 20, 30, 40, 50}
	s := array[:]
	fmt.Printf("切片s的长度：%d，容量：%d\n", len(s), cap(s))

	s1 := array[1:3]
	fmt.Printf("切片s1的长度：%d，容量：%d\n", len(s1), cap(s1))
}
