package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ccmorenov/microservicesounds/bd"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r)
	ID := r.URL.Query().Get("id")
	fmt.Println("ID: ", ID)
	if len(ID) < 1 {
		http.Error(w, "Falta el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.SearchPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro "+err.Error(), 400)
		return
	}

	token := r.URL.Query().Get("token")
	fmt.Println("token: ", token)
	_, _, _, err = ProcesoToken(token)

	if err != nil {
		http.Error(w, "Ocurrio un error con el token"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
