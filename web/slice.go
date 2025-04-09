package main

import (
    "log"
    "sort"
)

func main() {
    var mySlice []string
    mySlice = append(mySlice, "Joseph")
    mySlice = append(mySlice, "Jodie")
    mySlice = append(mySlice, "Stella")
    log.Println(mySlice)

    mySlice2 := []int{2, 1, 3, 7, 5, 4, 6, 9, 8}
    log.Println(mySlice2)
    sort.Ints(mySlice2)
    log.Println(mySlice2)
    log.Println(mySlice2[:3])
}
