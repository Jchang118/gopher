package main

import "fmt"

func main() {
	var count = 5
	msgChannel := make(chan int, count)

	for i := 0; i < count; i++ {
		msgChannel <- i
	}

	for i := 0; i < count; i++ {
		fmt.Println(<- msgChannel)
	}
}
