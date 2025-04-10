package main

import "log"

func main() {

    for i := 0; i <= 10; i++ {
        log.Println(i)
    }

    log.Println()

    animals := []string{"dog", "fish", "horse", "cow", "cat"}
    for _, animial := range animals {
        log.Println(animial)
    }

    log.Println()

    animals_map := make(map[string]string)
    animals_map["dog"] = "Fido"
    animals_map["cat"] = "Fluffy"

    for k, v := range animals_map {
        log.Println(k, v)
    }

    log.Println()

    var firstLine = "Once upon a midnight dreary"

    for i, l := range firstLine {
        log.Printf("%d : %c\n", i, l)
    }

    log.Println()

    type User struct {
        FirstName string
        LastName string
        Email string
        Age int
    }

    var users []User
    users = append(users, User{"John", "Smith", "john@smith.com", 30})
    users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
    users = append(users, User{"Sally", "Brown", "sally@brown.com", 45})
    users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 17})

    for _, l := range users {
        log.Println(l.FirstName, l.LastName, l.Email, l.Age)
    }
}
