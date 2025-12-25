package utils

import (
	"log"
	"net/http"
)

func LogRequest(req *http.Request) {
	log.Printf("%s %s %s %d\n", req.Method, req.URL, req.Proto, http.StatusOK)
}