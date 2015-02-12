package main

import (
    "fmt"
    "log"
    "time"
    "strings"
    "errors"
    "net/http"
    "encoding/json"
    "github.com/justinas/alice"
    "github.com/gorilla/context"
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

func getUser(auth string) (string, error) {
    if strings.Contains(auth, "Authorization") {
        return auth, nil
    } else {
        err := errors.New("Logon failed")
        return auth, err
    }
}

func authHandler(next http.Handler) http.Handler {
    fn := func(w http.ResponseWriter, r *http.Request) {
        authToken := r.Header.Get("Authorization")
        user, err := getUser(authToken)

        if err != nil {
            http.Error(w, http.StatusText(401), 401)
            return
        }

        context.Set(r, "user", user)
        next.ServeHTTP(w, r)
    }

    return http.HandlerFunc(fn)
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
    user := context.Get(r, "user")
    json.NewEncoder(w).Encode(user)
}

func main() {
    commonHandlers := alice.New(context.ClearHandler, loggingHandler, recoverHandler)
    http.Handle("/admin", commonHandlers.Append(authHandler).ThenFunc(adminHandler))
    http.Handle("/uve", commonHandlers.ThenFunc(uveHandler))
    http.Handle("/", commonHandlers.ThenFunc(indexHandler))
    http.ListenAndServe(":8169", nil)
}
