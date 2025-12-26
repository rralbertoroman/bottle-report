package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rralbertoroman/bottle-report/internal/app"
	"github.com/rralbertoroman/bottle-report/internal/middleware"
)

func Health(a *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := a.DB.PingContext(r.Context()); err != nil {
			http.Error(w, "db unavailable", http.StatusServiceUnavailable)
			return
		}
		w.Write([]byte("ok"))
	}
}

func MessagesHandler(a *app.App) http.HandlerFunc{
	return middleware.WithLogging(
		func (w http.ResponseWriter, req *http.Request) (status int){
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
		},
	)
}
