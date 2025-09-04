package api

import (
	"net/http"

	"github.com/gorilla/csrf"
	jsonutils "github.com/reinheimermat/gobid/internal/json_utils"
)

func (api *Api) handleGetCSRFToken(w http.ResponseWriter, r *http.Request) {
	token := csrf.Token(r)
	jsonutils.EncodeJSON(w, r, http.StatusOK, map[string]any{
		"csrf_token": token,
	})
}

func (api *Api) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			jsonutils.EncodeJSON(w, r, http.StatusUnauthorized, map[string]any{
				"message": "user not authenticated",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
