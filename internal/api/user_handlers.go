package api

import "net/http"

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("signup user"))
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login user"))
}

func (api *Api) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("logout user"))
}
