package main

import "fmt"

func main() {
	const (
		Sunday = iota
		Monday = iota
		Tuesday = iota
		Wednesday = iota
		Thursday = iota
		Friday = iota
		Saturday = iota
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
