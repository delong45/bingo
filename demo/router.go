package main

import (
    "net/http"
    "fmt"
    "time"
    "log"
    "github.com/gorilla/mux"
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

func idHandler(w http.ResponseWriter, r *http.Request) {
    t1 := time.Now()
    vars := mux.Vars(r)
    fmt.Fprintf(w, "Show your id %v", vars["id"])
    t2 := time.Now()
    log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/articles/{id:[0-9]+}", idHandler)
    router.HandleFunc("/uve", uveHandler)
    router.HandleFunc("/", indexHandler)
    http.Handle("/", router)
    http.ListenAndServe(":8191", nil)
}
