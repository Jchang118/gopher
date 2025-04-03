package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	s := array[1:3]
	for i, e := range s {
		fmt.Printf("第%d个元素为：%d\n", i, e)
	}
}
