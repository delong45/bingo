package main

import (
    "fmt"
    "log"
    "time"
    "net/http"
    "github.com/justinas/alice"
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

func loggingHandler(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        t1 := time.Now()
        next.ServeHTTP(w, r)
        t2 := time.Now()
        log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
    }

    return http.HandlerFunc(fn)
}

func recoverHandler(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("panic: %+v", err)
                http.Error(w, http.StatusText(500), 500)
            }
        }()

        next.ServeHTTP(w, r)
    }

    return http.HandlerFunc(fn)
}

func main() {
    commonHandlers := alice.New(loggingHandler, recoverHandler)
    http.Handle("/uve", commonHandlers.ThenFunc(uveHandler))
    http.Handle("/", commonHandlers.ThenFunc(indexHandler))
    http.ListenAndServe(":8169", nil)
}
