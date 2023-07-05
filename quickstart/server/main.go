package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user", userHandle)
	http.ListenAndServe(":9090", nil)
}

func userHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Fprintln(w, "hello")
}
