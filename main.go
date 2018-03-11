package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/qinains/fastergoding"
)

func main() {
	fastergoding.Run() // Just add this code

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %s!", r.URL.Query().Get("name"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
