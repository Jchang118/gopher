package main

import "fmt"

func main() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday = "x"
		Thursday = iota
		Friday = iota
		Saturday = iota
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
