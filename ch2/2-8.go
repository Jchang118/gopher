package main

import "fmt"

func main() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday = "x"
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
