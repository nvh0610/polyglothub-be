package flashcard_daily

import "net/http"

type Controller interface {
	GetFlashCardDaily(w http.ResponseWriter, r *http.Request)
	ConfirmFlashCardDaily(w http.ResponseWriter, r *http.Request)
}
