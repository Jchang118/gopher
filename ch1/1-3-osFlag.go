package main

import (
	"flag"
	"fmt"
)

func main() {
	var intVal = flag.Int("intVal", 0, "int类型参数")
	var boolVal = flag.Bool("boolVal", false, "bool类型参数")
	var stringVal = flag.String("stringVal", "", "string类型参数")

	flag.Parse()

	fmt.Println("-intVal: ", *intVal)
	fmt.Println("-boolVal", *boolVal)
	fmt.Println("-stringVal", *stringVal)
}
