package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "OK"
	StatusError = "Error"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(err validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range err {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, err.Field()+" is required")
		case "email":
			errMsgs = append(errMsgs, err.Field()+" must be a valid email")
		case "gte":
			errMsgs = append(errMsgs, err.Field()+" must be greater than or equal to "+err.Param())
		case "lte":
			errMsgs = append(errMsgs, err.Field()+" must be less than or equal to "+err.Param())
		default:
			errMsgs = append(errMsgs, err.Field()+" is not valid")
		}
	}
	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}

}
