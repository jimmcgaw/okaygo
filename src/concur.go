package main

import (
  "fmt"
  "time"
  "sync"
  "runtime"
)

func main() {
  
  // optional line to parallelize all of this
  runtime.GOMAXPROCS(2)
  
  var waitGroup sync.WaitGroup
  waitGroup.Add(2)
  
  go func() {
    defer waitGroup.Done()
    
    time.Sleep(5 * time.Second)
    fmt.Println("Hello")
  }()
  
  go func() {
    defer waitGroup.Done()
    fmt.Println("Goodbye")
  }()
  
  waitGroup.Wait()
}