package main

import (
    "fmt"
)

func celebrateBirthday(people map[string]int, name string) {
    people[name]++
}

func main() {
    myAgeMap := make(map[string]int)
    myAgeMap["Jim"] = 34
    myAgeMap["Rodrigo"] = 26
    fmt.Println(myAgeMap)
    
    // literal declaration syntax for maps
    mySalariesMap := map[string]float64 {
        "Jim": 0.0,
        "Rodrigo": 1000.00, 
    }
    fmt.Println(mySalariesMap)
    
    for key,value := range myAgeMap {
        fmt.Printf("%s tiene %d años.\n", key, value)
    }
    
    celebrateBirthday(myAgeMap, "Rodrigo")
    
    for key,value := range myAgeMap {
        fmt.Printf("%s tiene %d años.\n", key, value)
    }
    
    myAgeMap["Casper"] = 66
    
    fmt.Println(myAgeMap)
    delete(myAgeMap, "Casper")
    fmt.Println(myAgeMap)
    
    // if we're going to create a huge map, declare its capacity upfront
    myHugeMap := make(map[string]int, 600)
    fmt.Printf("Capacity of 'myHugeMap' is: %d", len(myHugeMap)) // can we check size?
    
}