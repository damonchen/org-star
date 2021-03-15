package view

import "net/http"

// Index index
func Index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello world"))
}
