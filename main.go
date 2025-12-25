package main

import (
	"log"
	"net/http"

	"github.com/rralbertoroman/bottle-report/messaging"
)

func main() {
	messaging.InitRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}