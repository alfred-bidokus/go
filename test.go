package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hallo")
}

func main() {
	port := os.Getenv("PORT")
	fmt.Println(port)
	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)
}
