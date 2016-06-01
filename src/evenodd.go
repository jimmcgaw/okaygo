package main

import (
    "fmt"
    "math/rand"
    "time"
)

func random() int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(10) // return int b/w 0 and 9
}

func main() {
    // assignment before switch blog is "simple statement"
    switch tmpNum := random(); tmpNum {
        case 0, 2, 4, 6, 8:
        fmt.Println("Got an even value of", tmpNum)
        case 1, 3, 5, 7, 9:
        fmt.Println("Got an odd value of", tmpNum)
    }
}