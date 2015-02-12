package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    times := 1000000
    start := time.Now()

    for i := 0; i < times; i++ {
        conn, err := net.Dial("tcp", "127.0.0.1:1234")
        if err != nil {
            fmt.Printf("Dial error: %s\n", err.Error())
            continue
        }
        fmt.Printf("%d\n", i)

        msg := fmt.Sprintf("Http Hello Bingo")
        _, err = conn.Write([]byte(msg))
        if err != nil {
            fmt.Printf("Write error: %s\n", err.Error())
            conn.Close()
            continue
        }
        conn.Close()
        time.Sleep(1 * time.Microsecond)
    }

    end := time.Now()
    fmt.Printf("took %v to run.\n", end.Sub(start))
}
