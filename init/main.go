package main

import (
	"fmt"
	"net/http"
)

func indexhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello web billionere is back")
}
func move(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "millionere's  billionere")
}
func main() {
	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/keepMoving", move)
	http.ListenAndServe(":8000", nil)
}
