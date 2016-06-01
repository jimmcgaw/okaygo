package main

import (
    "fmt"
)

func main() {
    courses := make([]string, 5, 10)
    
    fmt.Printf("Length is: %d. \nCapacity is: %d\n", len(courses), cap(courses))
    
    coursesNoCapacity := make([]string, 5)
    fmt.Printf("Length is: %d. \nCapacity is: %d\n", len(coursesNoCapacity), cap(coursesNoCapacity))
    
    coursesLiteral := []string{ "Puppet", "Chef", "Ruby"}
    fmt.Printf("Length is: %d. \nCapacity is: %d\n", len(coursesLiteral), cap(coursesLiteral))
    
    numbers := []int{ 1,2,3,4,5,6,7,8,9,10}
    fmt.Println(numbers)
    
    sliceOfSlice := numbers[3:6]
    fmt.Println(sliceOfSlice)
    
    beginningOfSlice := numbers[:4]
    fmt.Println(beginningOfSlice)
    
    numbers = append(numbers, 11, 12, 13)
    fmt.Println(numbers)
    fmt.Printf("New capacity of numbers is: %d\n", cap(numbers))
    
    // combining slices
    nums := []int{ 1,2,3,4,5 }
    moreNums := []int { 60,70,80,90,100}
    nums = append(nums, moreNums...)
    fmt.Println(nums)
    
    firstSlice := []float64{ 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7 }
    secondSlice := firstSlice[2:5]
    firstSlice[4] = 0.99
    fmt.Println(secondSlice)  // notice that last value has been changed
    // this is because 1st and 2nd slice both refer to values
    // in the same underlying array
}