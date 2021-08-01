package delivery

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"go-boilerplate/domain/user/usecase"
)

type Delivery interface {
	ListHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type DefaultDelivery struct {
	Usecase usecase.Usecase
}
