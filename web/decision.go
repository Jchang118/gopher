package main

import "log"

func main() {

//    isTrue := true
//
//    if isTrue {
//        log.Println("isTrue is", isTrue)
//    } else {
//        log.Println("isTrue is", isTrue)
//    }
//
//    cat := "cat"
//
//    if cat == "cat" {
//        log.Println("Cat is cat")
//    } else {
//        log.Println("Cat is not cat")
//    }
//
//    myNum := 100
//    isTrue2 := false
//
//    if myNum > 99 && !isTrue2 {
//        log.Println("myNum is greater than 99 and isTrue2 is set to true")
//    } else if myNum < 100 && isTrue2 {
//        log.Println("xxxxx")
//    } else if myNum == 101 || isTrue2 {
//        log.Println("yyyyy")
//    } else {
//        log.Println("zzzzz")
//    }

myVar := "cat"
    switch myVar {
    case "cat":
        log.Println("cat is set to cat")
    case "dog":
        log.Println("cat is set to dog")
    case "fish":
        log.Println("cat is set to fish")
    default:
        log.Println("cat is something else")
    }
}
