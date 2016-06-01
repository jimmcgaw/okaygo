package main

import (
  "fmt"
  "strings"
)

func main() {
  phrase := "These are the times that try men's souls.\n"
  words := strings.Split(phrase, " ")
  
  
  ch := make(chan string, len(words))
  
  for _, word := range words {
    ch <- word
  }
  
  // this closes the channel for the sender; the receiver can still
  // pull from it.
  close(ch)
  
//  for i := 0; i < len(words); i++ {
//    fmt.Print(<-ch + " ")  // <- is the 'receive operator'
//  }
  
  // one way to skin this cat:
//  for {
//    if msg, isChannelOpen := <- ch; isChannelOpen {
//      fmt.Print(msg + " ")
//    } else {
//      break
//    }
//  }
  
  // Go's shorthand for the above for-if loop operation
  for msg := range ch {
    fmt.Print(msg + " ")
  }
  
  // this would fail with 'panic: attempted send on closed channel'
  // ch <- "New Word"
}