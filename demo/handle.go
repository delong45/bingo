package main

import (
    "net/http"
    "fmt"
    "time"
    "log"
)

func uveHandler(w http.ResponseWriter, r *http.Request) {
    t1 := time.Now()
    fmt.Fprintf(w, "You are on the UVE page")
    t2 := time.Now()
    log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    t1 := time.Now()
    fmt.Fprintf(w, "Welcome to Bingo")
    t2 := time.Now()
    log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
}

func main() {
    http.HandleFunc("/uve", uveHandler)
    http.HandleFunc("/", indexHandler)
    http.ListenAndServe(":9080", nil)
}
