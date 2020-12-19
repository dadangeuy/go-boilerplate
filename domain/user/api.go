package user

import (
	"github.com/julienschmidt/httprouter"
	"go-boilerplate/model"
	"net/http"
)

type Delivery interface {
	ListHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

type Usecase interface {
	List() ([]model.User, error)
}

type Repository interface {
	Create(email string, password string) (model.User, error)
	FindAll() ([]model.User, error)
}
