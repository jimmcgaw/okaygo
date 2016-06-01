package main

import (
    "fmt"
)

func changeCourse(course string) string {
    course = "First Look: Native Docker Clustering"
    
    fmt.Println("\nTrying to change your course to", course)
    
    return course
}

func main() {
    name := "Jim"
    course := "Docker Deep Dive"
    
    fmt.Println("\nHi", name, ", you're currently watching", course)
    
    changeCourse(course)
    
    fmt.Println("\nYou are currently watching", course)
}