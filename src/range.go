package main

import (
    "fmt"
)

func main() {
    // courses is a slice data type
    courses := []string{ "Go Fundamentals", "Distributed Programming in Go", "Node.js Etc" }
    for _, val := range courses {
        fmt.Println(val)
    }
}