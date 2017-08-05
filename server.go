package main

import (
	"fmt"
	"net/http"
	"os"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK\n")
}

func main() {
	http.HandleFunc("/healthz", healthzHandler)
	fmt.Fprint(os.Stdout, "Listening on :8424")
	http.ListenAndServe(":8424", nil)
}
