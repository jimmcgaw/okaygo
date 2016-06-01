package main

import (
    "fmt"
    "strings"
)

func converter(name, title string) (s1, s2 string) {
    name = strings.Title(name)
    title = strings.ToUpper(title)
    return name, title
}

func main() {
    name := "jim mcgaw"
    title := "software engineer"
    
    name, title = converter(name, title)
    fmt.Println("Converted name is", name)
    fmt.Println("Converted title is", title)
    
}