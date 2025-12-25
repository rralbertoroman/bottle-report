package messaging

import (
	"io"
	"net/http"

	"github.com/rralbertoroman/bottle-report/middleware"
)

func helloHandler(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "Hello world \n")
}

func InitRoutes() {
	http.HandleFunc("/hello", middleware.WithLogging(helloHandler))
}