
package main

import (
	"fmt"
	"log"
	"net/http"
	"webart"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../templates"))
	fmt.Print(fs)
	
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", webart.Handler)
	mux.HandleFunc("/ascii-art", webart.Asciiart)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
