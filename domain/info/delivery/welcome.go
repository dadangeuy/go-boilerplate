package delivery

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (h *Delivery) WelcomeHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "it works!")
}
