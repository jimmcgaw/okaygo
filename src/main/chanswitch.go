package main

import (
  "fmt"
)

func main() {
  msgCh := make(chan Message, 1)
  errCh := make(chan FailedMessage, 1)
}

type Message struct {
  To []string
  From string
  Content string
}

type FailedMessage struct {
  ErrorMessage string
  OriginalMessage string
}