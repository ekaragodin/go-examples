package main

import (
  "net/http"
  "fmt"
  "log"
  "io/ioutil"
)

var root = "/";

func main() {
  http.HandleFunc("/", indexHandler)
  log.Println("Server is started...")
  http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  entries, _ := ioutil.ReadDir(root)
  for _, entry := range entries {
    fmt.Println(entry.Name())
  }

  fmt.Fprintf(w, "Hello");
}
