package main

import "fmt"

func main() {
	charCountMap := make(map[string]int)
	count, ok := charCountMap["a"]
	fmt.Println("count:", count, ", exists:", ok)
}
