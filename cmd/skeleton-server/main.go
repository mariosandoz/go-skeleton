package main

import (
	"fmt"
	"go-skeleton/pkg/adding"
	"go-skeleton/pkg/http/rest"
	"log"
	"net/http"
)

func main() {
	var adder adding.Service
	router := rest.Handler(adder)
	fmt.Println("The creator server is on craft now: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
