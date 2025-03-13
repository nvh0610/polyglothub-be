package category

import "net/http"

type Controller interface {
	CreateCategory(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	ListCategory(w http.ResponseWriter, r *http.Request)
}
