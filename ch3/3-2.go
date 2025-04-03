package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := 3.14
	b := 3.15

	// 利用big.NewFloat()创建Float指针，两个Float指针进行比较
	result := big.NewFloat(a).Cmp(big.NewFloat(b))

	if result < 0 {
		fmt.Println("a小于b")
	}
	if result == 0 {
		fmt.Println("a等于b")
	}
	if result > 0 {
		fmt.Println("a大于b")
	}
}
