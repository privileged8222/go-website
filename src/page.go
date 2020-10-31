package main

import (
  "net/http"
)

//struct represents a page with a filename
type page_t struct {
  m_name string
}

//this function displays the page
func (p * page_t) Execute(_writer http.ResponseWriter, _request *http.Request) () {
  g_tpl.ExecuteTemplate(_writer, p.m_name, nil);
}

//we use this as a wrapper around the execute function
//to make it easier to call
func (p * page_t) Display(_path string) () {
  http.HandleFunc(_path, p.Execute) // pass in execute function
}
