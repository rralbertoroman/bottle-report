package middleware

import (
	"net/http"

	"github.com/rralbertoroman/bottle-report/internal/utils"
)

func WithLogging(next func(w http.ResponseWriter, r *http.Request ) (status int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := next(w, r)
		utils.LogRequest(r,status)
	}
}