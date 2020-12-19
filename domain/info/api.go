package info

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Delivery interface {
	WelcomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
