package main

import "fmt"

func main() {

	s := make([]int, 5, 10)

	fmt.Println(s)

	s = append(s, 5, 6, 7, 8, 9)

	fmt.Printf("追加后切片长度：%d\n", len(s))
	fmt.Printf("追加后切片容量：%d\n", cap(s))
	fmt.Println(s)
}
