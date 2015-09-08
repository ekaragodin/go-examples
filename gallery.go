package main

import (
  "net/http"
  "html/template"
  "log"
)

var templates = template.Must(template.ParseFiles("templates/gallery.html"))

func main() {
  http.HandleFunc("/", indexHandler)
  log.Println("Server is started...")
  http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  err := templates.ExecuteTemplate(w, "gallery.html", nil)
  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      log.Println(err.Error())
  }
}
