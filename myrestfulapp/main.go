// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/hello", HelloHandler).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Hello, World! i am uday"}`)
}
