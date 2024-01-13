package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	http.ListenAndServe(":3000", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Shell:", os.Getenv("IZIMONEY_DB_HOST"))

	fmt.Fprintf(w, "Hello World")
}
