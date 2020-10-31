package main

import (
  "fmt"
  "html/template"
  "net/http"
)

//var to represent the template pointer
var g_tpl *template.Template

func init() { // parse our pages
  g_tpl = template.Must(template.ParseGlob("../site/*.html"))
}

func main() {
  _port := ":8080" // port to work on
  fmt.Printf("Ready on port%s\n", _port)

  _about := page_t{"about.html"} // create about page

  //serve the website from out filesystem
  http.Handle("/", http.FileServer(http.Dir("../site/")))
  //display pages
  _about.Display("about/")

  http.ListenAndServe(_port, nil) // listen!
}
