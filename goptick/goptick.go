package goptick

import (
    "fmt"
    "time"
)

type fn func(chan bool)


func periodicTick(q chan bool, f fn, tick time.Time){
        fmt.Println("Tick at: ", tick)
        f(q)
}



func Gotick(d time.Duration, f fn){
    ticker := time.NewTicker(d)
    quit := make(chan bool, 1)

    go func() {
        for {
           select {
            case <- ticker.C:
                periodicTick(quit, f, <- ticker.C)
            }
        }
     }()

    <-quit

}
