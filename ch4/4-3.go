package main

import "fmt"

func main() {
	s := make([]int, 5, 10)

	for i := 0; i < len(s); i++ {
		s[i] = i
	}

	for _, e := range s {
		fmt.Println(e)
	}
}
