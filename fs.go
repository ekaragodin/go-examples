package main

import (
  "net/http"
  "log"
  "io/ioutil"
  "html/template"
)

var root = "/";

func main() {
  http.HandleFunc("/", indexHandler)
  log.Println("Server is started...")
  http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/fs.html")
  if err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }

  entries, _ := ioutil.ReadDir(root)
  t.Execute(w, entries)
}
