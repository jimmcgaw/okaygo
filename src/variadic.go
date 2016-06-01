package main

import (
    "fmt"
)

// numbers is a slice of ints
func getMaxNumber(numbers ...int) int {
    highest := numbers[0]  // first value by default
    for _, value := range numbers {
        if value > highest {
            highest = value
        }
    }
    return highest
}

func main() {
    maxNumber := getMaxNumber(1,2,3,6,7,4,99)
    fmt.Println("Max number is", maxNumber)
}