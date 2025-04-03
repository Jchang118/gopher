package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	s1 := s[1:]
	s1[0] = 0
	fmt.Println("s = ", s)
	fmt.Println("s1 = ", s1)
}
