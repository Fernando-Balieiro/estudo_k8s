package main

import (
	"net/http"
)

func main() {
  http.HandleFunc("/", hello)
  http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("<h1>Hello world from h8s!</h1>"))
}
