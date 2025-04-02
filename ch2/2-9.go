package main

import "fmt"

func main() {
	const (
		_ = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)

	fmt.Println( Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
}
