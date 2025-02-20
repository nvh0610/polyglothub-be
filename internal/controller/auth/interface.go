package auth

import "net/http"

type Controller interface {
	Login(w http.ResponseWriter, r *http.Request)
	ChangePassword(w http.ResponseWriter, r *http.Request)
}
