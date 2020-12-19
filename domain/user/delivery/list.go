package delivery

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (d *Delivery) ListHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	users, err := d.Usecase.List()
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		panic(err)
	}
}
