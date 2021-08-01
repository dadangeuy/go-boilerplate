package delivery

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *DefaultDelivery) WelcomeHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "it works!")
}
