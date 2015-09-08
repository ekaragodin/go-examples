package main

import (
  "net/http"
  "html/template"
  "log"
  "os"
  "path"
)

var templates = template.Must(template.ParseFiles("templates/gallery.html"))
var dataDir string

func main() {
  dataDir = getDataDir()

  http.HandleFunc("/", indexHandler)
  log.Println("Server is started...")
  http.ListenAndServe(":8000", nil)
}

func getDataDir() string {
  pwd, _ := os.Getwd()
  return path.Join(pwd, "data", "gallery")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  err := templates.ExecuteTemplate(w, "gallery.html", nil)
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      log.Println(err.Error())
  }
}
