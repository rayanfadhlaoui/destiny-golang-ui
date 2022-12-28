package main

import (
	"github.com/rayanfadhlaoui/destiny-golang-ui/pkg/handlers"
	"log"
	"net/http"
)

var repository *handlers.Repository

const portNumber = ":80"

func init() {
	repository = handlers.NewRepository(false)
}

func main() {
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(repository),
	}
	log.Println("started in port", portNumber)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
