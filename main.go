package main

import (
  "os"
  "fmt"
  "html"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
  // Using net/http
  //http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  //	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  //})
  //log.Fatal(http.ListenAndServe(":" + os.Getenv("HTTP_PLATFORM_PORT"), nil))

  // Using gorilla/mux
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Index)
  log.Fatal(http.ListenAndServe(":" + os.Getenv("HTTP_PLATFORM_PORT"), router))
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
