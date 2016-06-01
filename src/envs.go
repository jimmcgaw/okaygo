package main

import (
    "fmt"
    "os"
    "reflect"
)

func main() {
    // spoiler alert: it's a byte string
    fmt.Println("Variable type of environment variable hash is", reflect.TypeOf(os.Environ()))
    
    for _, env := range os.Environ() {
        fmt.Println(env)
    }
    
    username := os.Getenv("USER")
    
    fmt.Println("Name of current system user:", username)
}