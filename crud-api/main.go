package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello!")
	r := mux.NewRouter()
	fmt.Println(r)
}
