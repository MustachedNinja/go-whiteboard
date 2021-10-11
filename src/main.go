package main

import (
	"net/http"
	"whiteboard/src/handler"
)

func main() {
	http.HandleFunc("/helloWorld", handler.HelloWorld)
	http.ListenAndServe(":5000", nil)
}
