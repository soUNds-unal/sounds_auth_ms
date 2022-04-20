package routes

import (
	"net/http"

	"github.com/ccmorenov/microservicesounds/bd"
)

func EliminarUser(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar parametro de ID", http.StatusBadRequest)
		return
	}

	err := bd.BorroUser(IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar al usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
