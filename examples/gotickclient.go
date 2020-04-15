package main

import (
    "fmt"
    "time"
    tic "github.com/goprag/gotick/goptick"
    )

var count int
func foo(q chan bool){
   fmt.Println("bar baz...")
   if count > 10 {
       q <- true
   }
   count++
}

func main() {
    count = 0
    d := 1*time.Second
    tic.Gotick(1*d, foo)
    fmt.Println("done")

}