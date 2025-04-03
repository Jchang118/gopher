package main

import "fmt"

func main() {
	s := "字符串"

	len := len(s)
	fmt.Println(len)
	
	for i := 0; i < len; i++ {
		fmt.Printf("%X ", s[i])
	}
}
