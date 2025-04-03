package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "字符串"
	// 利用utf8.RuneCountInString()代替len()函数，来获得字符数据量
	fmt.Println("字符串的长度为：", utf8.RuneCountInString(s))

	// 利用for-range循环代替s[i]，来获得单个字符
	for _, c := range s{
		fmt.Printf("%T, %X\n", c, c)
	}
}
