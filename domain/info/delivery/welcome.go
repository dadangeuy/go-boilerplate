package delivery

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (h *Delivery) WelcomeHandler(writer http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	writer.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(writer, "it works!")
}
