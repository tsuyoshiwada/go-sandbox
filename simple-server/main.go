package main

import (
  "fmt"
  "net/http"
  "strings"
  "log"
)

const Port int = 8080

func sayhelloName(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Println(r.Form)
  fmt.Println("Path:", r.URL.Path)
  fmt.Println("Scheme: ", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])
  for k, v := range r.Form {
    fmt.Println("key:", k)
    fmt.Println("val:", strings.Join(v, ""))
  }
  fmt.Fprintf(w, "Hello world!!")
}

func main() {
  http.HandleFunc("/", sayhelloName)
  err := http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
  if err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
