package messaging

import (
	"encoding/json"
	"net/http"

	"github.com/rralbertoroman/bottle-report/middleware"
)

func MessagesHandler(w http.ResponseWriter, req *http.Request) (status int){
	switch req.Method {
	case http.MethodPost:
		SaveMessage(req.Body)
	case http.MethodGet:
		messages := AllMessages()

		if len(messages) == 0 {
			status = http.StatusNotFound
			w.WriteHeader(status)
			return
		}

		json.NewEncoder(w).Encode(messages)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		status =  http.StatusMethodNotAllowed
		return
	}

	status = http.StatusOK
	return
}

func InitRoutes() {
	http.HandleFunc("/messages", middleware.WithLogging(MessagesHandler))
}