package main

import (
    "fmt"
    "reflect"
)

func main() {
    a := 10
    b := 3.0000000
    
    fmt.Println("\na is of type:", reflect.TypeOf(a), "and b is of type:", reflect.TypeOf(b));
    
    c := float64(a) + b
    
    fmt.Println("\nC has value:", c, "and is of type:", reflect.TypeOf(c))
}