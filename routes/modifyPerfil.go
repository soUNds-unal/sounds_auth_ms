package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ccmorenov/microservicesounds/bd"
	"github.com/ccmorenov/microservicesounds/models"
)

func ModifyPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.ModifyRegister(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error, intente modificar de nuevo "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro de usuario "+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
