package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	fmt.Printf("切片的长度：%d\n", len(s))
	fmt.Printf("切片的容量：%d\n", cap(s))
}
