package delivery

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Delivery interface {
	WelcomeHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type DefaultDelivery struct{}
