package main

import (
  "fmt"
)

func main() {
  
}

type Button struct {
  eventListeners map[string][]chan string
  
}

func MakeButton() *Button {
  result := new(Button)
  result.eventListeners = make(map[string][]chan string)
  return result
}

func (this *Button) AddEventListener(event string, responseChannel chan string) {
  if _, isPresent := this.eventListeners[event]; isPresent {
    this.eventListeners[event] = append(this.eventListeners[event], responseChannel)
  } else {
    this.eventListeners[event] = []chan string{responseChannel}
  }
}

func (this *Button) RemoveEventListener(event string, listenerChannel chan string) {
  if _, isPresent := this.eventListeners[event]; isPresent {
    for idx, _ := range this.eventListeners[event] {
      if this.eventListeners[event][idx] == listenerChannel {
        this.eventListeners[event] = append(this.eventListeners[event][:idx], 
                                           this.eventListeners[event][idx+1]...)
        break
      }
    }
  }
}