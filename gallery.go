package main

import (
  "net/http"
  "html/template"
  "log"
  "os"
  "path"
  "io/ioutil"
  "strings"
)

var templates = template.Must(template.ParseFiles("templates/gallery.html"))
var dataDir string

type Image struct {
  Src string
}

func main() {
  dataDir = getDataDir()

  static := http.FileServer(http.Dir(dataDir))
  http.Handle("/images/", http.StripPrefix("/images/", static))
  http.HandleFunc("/", indexHandler)

  log.Println("Server is started...")
  http.ListenAndServe(":8000", nil)
}

func getDataDir() string {
  pwd, _ := os.Getwd()
  return path.Join(pwd, "data", "gallery")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
  err := templates.ExecuteTemplate(w, "gallery.html", getImages())

  if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      log.Println(err.Error())
  }
}

func getImages() []Image {
  images := []Image{}
  files, _ := ioutil.ReadDir(dataDir)
  for _, file := range files {
    if (strings.HasPrefix(file.Name(), ".")) {
      continue
    }

    image := Image{
      Src: file.Name(),
    }
    images = append(images, image)
  }

  return images
}
