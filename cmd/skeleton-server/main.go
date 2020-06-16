package main

import (
	"fmt"
	"go-skeleton/pkg/adding"
	"go-skeleton/pkg/http/rest"
	"go-skeleton/pkg/storage/mongodb"
	"log"
	"net/http"
)

func main() {

	r, _ := mongodb.NewStorage()
	adder := adding.NewService(r)

	// set up the HTTP server
	router := rest.Handler(adder)

	fmt.Println("The creator server is on craft now: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
