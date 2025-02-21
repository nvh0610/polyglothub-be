package vocabulary

import "net/http"

type Controller interface {
	CreateVocabulary(w http.ResponseWriter, r *http.Request)
	UpdateVocabulary(w http.ResponseWriter, r *http.Request)
	DeleteVocabulary(w http.ResponseWriter, r *http.Request)
	ListVocabulary(w http.ResponseWriter, r *http.Request)
}
