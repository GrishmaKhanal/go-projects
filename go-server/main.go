package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse Form() err: %v", err)
	}
	fmt.Fprintf(w, "POST request sucessfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "\nName = %v\nAddress = %v", name, address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting the server in port: 8080\nurl: http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
}
