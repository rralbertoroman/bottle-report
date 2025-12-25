package messaging

import (
	"net/http"

	"github.com/rralbertoroman/bottle-report/middleware"
)

func MessageHandler(w http.ResponseWriter, req *http.Request){
	switch req.Method {
	case http.MethodPost:
		SaveMessage(req.Body)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func InitRoutes() {
	http.HandleFunc("/messages", middleware.WithLogging(MessageHandler))
}