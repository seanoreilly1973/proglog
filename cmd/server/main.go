package main

import (
	"log"

	"github.com/seanoreilly1973/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
