package main

import (
  "net/http"
  "log"
  "io/ioutil"
  "html/template"
  "os"
  "sort"
  "path"
)

var root = "/";

type Entry struct {
  Name string
  FullName string
  IsDir bool
}

type ByIsDir []Entry

func (slice ByIsDir) Len() int {
  return len(slice)
}

func (slice ByIsDir) Less(i, j int) bool {
  if (slice[i].IsDir != slice[j].IsDir) {
    return slice[i].IsDir
  }

  return slice[i].Name < slice[j].Name
}

func (slice ByIsDir) Swap(i, j int) {
  slice[i], slice[j] = slice[j], slice[i]
}

func main() {
  http.HandleFunc("/", indexHandler)
  log.Println("Server is started...")
  http.ListenAndServe(":8000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  currentPath := r.URL.Query().Get("path")
  if (currentPath == "") {
    currentPath = root
  }

  if _, err := os.Stat(currentPath); os.IsNotExist(err) {
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

  t.Execute(w, getEntries(currentPath))
}

func getEntries(currentPath string) []Entry {
  entries := []Entry{}

  if (currentPath != root) {
    parent := Entry{
      Name: "..",
      FullName: path.Join(currentPath, ".."),
      IsDir: true,
    }
    entries = append(entries, parent)
  }

  files, _ := ioutil.ReadDir(currentPath)
  for _, e := range files {
    entries = append(entries, Entry{
      Name: e.Name(),
      FullName: path.Join(currentPath, e.Name()),
      IsDir: e.IsDir(),
    })
  }

  sort.Sort(ByIsDir(entries))
  return entries
}
