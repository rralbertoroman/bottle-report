package middleware

import (
	"net/http"

	"github.com/rralbertoroman/bottle-report/utils"
)

func WithLogging(next func(w http.ResponseWriter, r *http.Request )) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		utils.LogRequest(r)
	}
}