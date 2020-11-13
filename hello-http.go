package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
)

func hello(w http.ResponseWriter, r *http.Request) {
    helloEnv := os.Getenv("HELLO")
    
    if helloEnv == "" {
        helloEnv = "hello"
    }

    arg := "class"
    if len(os.Args) > 1 {
        arg = os.Args[1]
    }

    log.Printf("%s %s\n", helloEnv, arg)
    
    fmt.Fprintf(w, "%s %s\n", helloEnv, arg)
}

func main() {

    http.HandleFunc("/hello", hello)
    log.Println("Starting server on 0.0.0.0:8080")
    http.ListenAndServe(":8080", nil)
}