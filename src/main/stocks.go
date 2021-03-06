package main

import (
  "net/http"
  "io/ioutil"
  "encoding/xml"
  "fmt"
  "time"
  "runtime"
)

func main() {
  start := time.Now()
  runtime.GOMAXPROCS(2)
  
  stockSymbols := []string {
    "googl",
    "msft",
    "aapl",
    "bbry",
    "hpq",
  }
  
  numComplete := 0
  
  for _, symbol := range stockSymbols {
    
    go func(symbol string){
      
      resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + symbol)
      defer resp.Body.Close()
      body, _ := ioutil.ReadAll(resp.Body)

      quote := new(QuoteResponse)
      xml.Unmarshal(body, &quote)
      fmt.Printf("%s: %.2f\n", quote.Name, quote.LastPrice)
      
      numComplete++
      
    }(symbol)

  }
  
  for numComplete < len(stockSymbols) {
    time.Sleep(time.Millisecond * 10)
  }
  
  fmt.Printf("Completed %d calls", numComplete)
  
  
  elapsed := time.Since(start)
  fmt.Printf("Execution time: %s", elapsed)
}


type QuoteResponse struct {
  Status string
  Name string
  LastPrice float32
  Change float32
  ChangePercent float32
  TimeStamp string
  MSDate float32
  MarketCap int
  Volume int
  ChangeYTD float32
  ChangePercentageYTD float32
  High float32
  Low float32
  Open float32
}