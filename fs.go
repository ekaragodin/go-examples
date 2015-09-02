package main

import (
  "net/http"
  "fmt"
  "log"
)

func main() {
  http.HandleFunc("/", indexHandler)
  log.Println("Server is started...")
  http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello");
}
