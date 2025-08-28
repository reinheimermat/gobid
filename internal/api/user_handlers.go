package api

import (
	"errors"
	"net/http"

	jsonutils "github.com/reinheimermat/gobid/internal/json_utils"
	"github.com/reinheimermat/gobid/internal/services"
	"github.com/reinheimermat/gobid/internal/usecase/user"
)

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJSON[user.CreateUserReq](r)

	if err != nil {
		_ = jsonutils.EncodeJSON(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.CreateUser(
		r.Context(),
		data.UserName,
		data.Email,
		data.Password,
		data.Bio,
	)

	if err != nil {
		errors.Is(err, services.ErrDuplicatedEmailOrUsername)
		_ = jsonutils.EncodeJSON(w, r, http.StatusUnprocessableEntity, map[string]any{
			"error": "email or username already exists",
		})
		return
	}

	_ = jsonutils.EncodeJSON(w, r, http.StatusOK, map[string]any{
		"user_id": id,
	})
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login user"))
}

func (api *Api) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout user"))
}
