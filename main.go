package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rralbertoroman/bottle-report/messaging"
)

func main() {
	messaging.InitRoutes()
	fmt.Println("Bottle Report")
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}