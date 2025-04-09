package main

import "log"

type User struct {
    FirstName string
    LastName string
}

func main() {

    myMap := make(map[string]string)

    myMap["dog"] = "Samson"
    myMap["other-dog"] = "Cassie"
    myMap["dog"] = "fido"

    log.Println(myMap["dog"])
    log.Println(myMap["other-dog"])

    myMap2 := make(map[string]int)
    myMap2["First"] = 1
    myMap2["Second"] = 2
    log.Println(myMap2["First"], myMap2["Second"])

    myMap3 := make(map[string]User)
    me := User {
        FirstName: "Joseph",
        LastName: "Chang",
    }

    myMap3["me"] = me
    log.Println(myMap3["me"].FirstName)

}
