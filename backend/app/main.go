package main

// import (
//   "log"
//   "net/http"
//   "math/rand"
//   "strconv"
//   "encoding/json"
//   "github.com/gorilla/mux"
// )

import (
    "net/http"
    "os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Hello World!</h1>"))
}

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    // env := os.Getenv("APP_ENV")

    mux := http.NewServeMux()

    mux.HandleFunc("/", indexHandler)
    http.ListenAndServe(":"+port, mux)
}
