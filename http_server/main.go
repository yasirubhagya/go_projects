package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		print(err.Error())
	}
}
