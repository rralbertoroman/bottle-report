package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rralbertoroman/bottle-report/internal/app"
	"github.com/rralbertoroman/bottle-report/internal/middleware"
	"gorm.io/gorm"
)

func Health(a *app.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		sqlDB, err := a.DB.DB()
		if err != nil {
			http.Error(w, "db unavailable", http.StatusServiceUnavailable)
			return
		}
		if err := sqlDB.Ping(); err != nil {
			http.Error(w, "db unavailable", http.StatusServiceUnavailable)
			return
		}
		w.Write([]byte("ok"))
	}
}

func MessagesHandler(a *app.App) http.HandlerFunc{
	a.DB.AutoMigrate(&Message{})
	ctx := context.Background()

	return middleware.WithLogging(
		func (w http.ResponseWriter, req *http.Request) (status int){
			var err error
			
			switch req.Method {
			case http.MethodPost:
				err = SaveMessage(ctx, req.Body, a)

			case http.MethodGet:
				var messages []Message
				messages, err = AllMessages(ctx, a)

				if len(messages) == 0 {
					status = http.StatusNotFound
					w.WriteHeader(status)
					return
				}

				json.NewEncoder(w).Encode(messages)
			case http.MethodDelete:
				id := req.URL.Query().Get("id")
				
				if id == "" {
					w.WriteHeader(http.StatusBadRequest)
					status = http.StatusBadRequest
					return
				}

				err = DeleteMessage(ctx, id, a)

				if err == gorm.ErrRecordNotFound {
					status = http.StatusNotFound
					w.WriteHeader(status)
					w.Write([]byte("Message not found"))
					return
				}
			default:
				w.WriteHeader(http.StatusMethodNotAllowed)
				status = http.StatusMethodNotAllowed
				return
			}

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}

			status = http.StatusOK
			return
		},
	)
}
