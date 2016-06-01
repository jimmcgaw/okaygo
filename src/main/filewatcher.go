package main

import (
  "fmt"
  "os"
  "time"
  "io/ioutil"
  "strings"
  "encoding/csv"
  "strconv"
  "runtime"
)

const watchedPath = "./source"

func main() {
  
  runtime.GOMAXPROCS(4)
  
  for {
    d, _ := os.Open(watchedPath)
    files, _ := d.Readdir(-1)
    for _, fi := range files {
      filePath := watchedPath + "/" + fi.Name()
      f, _ := os.Open(filePath)
      data, _ := ioutil.ReadAll(f)
      f.Close()  // note the lack of 'defer' here.
      os.Remove(filePath)
      
      go func(data string){
        reader := csv.NewReader(strings.NewReader(data))
        records, _ := reader.ReadAll();
        for _, record := range records {
          invoice := new(Invoice)
          invoice.Number = record[0]
          invoice.Amount, _ = strconv.ParseFloat(record[1], 64)
          invoice.PurchaseOrderNumber, _ = strconv.Atoi(record[2])
          unixTime, _ := strconv.ParseInt(record[3], 10, 64)
          invoice.InvoiceDate = time.Unix(unixTime, 0)
          
          fmt.Printf("Received invoice %v for $%.2f\n", invoice.Number, invoice.Amount)
        }
      }(string(data))
    }
  }
}

type Invoice struct {
  Number string
  Amount float64
  PurchaseOrderNumber int
  InvoiceDate time.Time
}