package main

import (
    "log"
    "time"
)

var s = "seven"

type User struct {
    FirstName   string
    LastName    string
    PhoneNumber string
    Age         int
    BirthDate   time.Time
}

func main() {
    user := User{FirstName: "Joseph", 
                LastName: "Chang",
                PhoneNumber: "0435 826 015",}
    log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate, user.PhoneNumber)
}

