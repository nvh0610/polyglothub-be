package user

import "net/http"

type Controller interface {
	GetUserById(w http.ResponseWriter, r *http.Request)
}
