package api

import (
	"net/http"

	jsonutils "github.com/reinheimermat/gobid/internal/json_utils"
)

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
