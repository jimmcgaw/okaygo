package main

import (
    "fmt"
)

//var (
//    //    name string
//    //    course string
//    name, course string
//    module float64
//)
//

// these would be globals
//var name, course, module = "Jim", "Go Fundamentals", 2.6


func main() {
    name := "Jim McGaw"
    course := "Go Fundamentals"
    module := 3.2
    
    fmt.Println("Name is set to: ", name)
    fmt.Println("Course is set to: ", course)
    fmt.Println("Module is set to: ", module)
}