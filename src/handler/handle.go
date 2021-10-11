package handler

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
