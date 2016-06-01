package main

import (
    "fmt"
)

func main() {
    module := 6.6
    
    fmt.Println("Module is", module)
    
    fmt.Println("Memory address of 'module' is:", &module)
    
    p := &module
    fmt.Println("Value of 'p' is:", *p)
}