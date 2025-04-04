package main

import "fmt"

type Person struct {
	Name	string
	Age		int
	Email	string
}

func main() {
	var person = Person {
		Name: "张三",
		Age: 20,
		Email: "78432@xxx.com",
	}
	
	fmt.Println("person: ", person)
	
	var person1 = Person{}
	fmt.Printf("person1 %T\n", person1)

	var person2 = new(Person)
	fmt.Printf("person2 %T\n", person2)
}
