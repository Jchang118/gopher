package main

import "fmt"

func main() {
	charCountMap := make(map[string]int)

	charCountMap["a"] = 3
	charCountMap["b"] = 5

	for k, v := range charCountMap {
		fmt.Printf("char:%s, count:%d\n", k, v)
	}
}
