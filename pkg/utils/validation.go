package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func BindAndValidate(r *http.Request, req interface{}) error {
	body := json.NewDecoder(r.Body)
	body.DisallowUnknownFields()

	if err := body.Decode(req); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return err
	}

	return nil
}
