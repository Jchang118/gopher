package main

import (
    "log"
    "package_example/helpers"
)

func main() {
    log.Println("Hello")

    var myVar helpers.SomeType
    myVar.TypeName = "Some Name"
    log.Println(myVar.TypeName)
}
