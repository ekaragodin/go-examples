package main

import (
  "net/http"
  "log"
  "io/ioutil"
  "html/template"
  "os"
)

var root = "/";

type Entry struct {
  Name string
  FullName string
  IsDir bool
}

func main() {
  http.HandleFunc("/", indexHandler)
  log.Println("Server is started...")
  http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Query().Get("path")
  if (path == "") {
    path = root
  }

  if _, err := os.Stat(path); os.IsNotExist(err) {
    log.Println(err.Error())
    http.Error(w, http.StatusText(404), 404)
    return
  }

  t, err := template.ParseFiles("templates/fs.html")
  if err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
    return
  }

  t.Execute(w, getEntries(path))
}

func getEntries(path string) []Entry {
  entries := []Entry{}
  files, _ := ioutil.ReadDir(path)
  for _, e := range files {
    entries = append(entries, Entry{
      Name: e.Name(),
      FullName: path + "/" + e.Name(),
      IsDir: e.IsDir(),
    })
  }
  return entries
}
