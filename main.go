package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Autor Łukasz Przychodzień - nasłuchuję na 8084")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})
	log.Fatal(http.ListenAndServe(":8084", nil))
}
