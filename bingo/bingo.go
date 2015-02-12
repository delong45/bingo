package main

import (
    "fmt"
    "net"
    "io"
    "os"
    "time"
)

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

func main() {
    listener, err := net.Listen("tcp", "0.0.0.0:1234")
    checkError(err)

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Fprintf(os.Stderr, "Accept error: %s", err.Error())
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()

    buf := make([]byte, 1024)
    for {
        n, err := conn.Read(buf)
        switch err {
        case nil:
            processData(buf[0:n], conn)
        case io.EOF:
            //fmt.Printf("Warnning: end of data: %s\n", err.Error)
            return
        default:
            fmt.Printf("Error: Reading data: %s\n", err.Error())
            return
        }
    }
}

func processData(buf []byte, conn net.Conn) {
    time.Sleep(50 * time.Millisecond)
}
