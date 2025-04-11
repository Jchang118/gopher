package main

import (
	"fmt"
	"net/http"
)

func main() {
	// fmt.Println("Hello World")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
		fmt.Println("Number of bytes written:", n)
	})

	_ = http.ListenAndServe(":8080", nil)
}
