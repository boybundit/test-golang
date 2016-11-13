package main

import (
  "os"
  "fmt"
  "html"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "time"
)

type Todo struct {
  Name      string
  Completed bool
  Due       time.Time
}

type Todos []Todo

func main() {
  // Using net/http
  //http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  //	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  //})
  //log.Fatal(http.ListenAndServe(":" + os.Getenv("HTTP_PLATFORM_PORT"), nil))

  // Using gorilla/mux
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Index)
  router.HandleFunc("/todos", TodoIndex)
  router.HandleFunc("/todos/{todoId}", TodoShow)

  log.Fatal(http.ListenAndServe(":" + os.Getenv("HTTP_PLATFORM_PORT"), router))
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  todos := Todos{
    Todo{Name: "Write presentation"},
    Todo{Name: "Host meetup"},
  }
  json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  todoId := vars["todoId"]
  fmt.Fprintln(w, "Todo show:", todoId)
}
