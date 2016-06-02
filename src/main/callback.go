package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main(){
  po := new(PurchaseOrder)
  
  po.Value = 69.99
  
  ch := make(chan *PurchaseOrder)
  
  go SavePO(po, ch)
  
  newPo := <- ch
  fmt.Printf("PO: %v", newPo)
  
}

type PurchaseOrder struct {
  Number int
  Value float64
}

func random() int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(10000) // return int b/w 0 and 9
}

func SavePO(po *PurchaseOrder, callback chan *PurchaseOrder) {
  po.Number = random()
  
  callback <- po
}