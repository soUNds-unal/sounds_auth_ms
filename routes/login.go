package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ccmorenov/microservicesounds/bd"
	"github.com/ccmorenov/microservicesounds/jwt"
	"github.com/ccmorenov/microservicesounds/middlew"
	"github.com/ccmorenov/microservicesounds/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y contraseña invalidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "EL mail del usuario es requerido ", 400)
		return
	}
	documento, existe := bd.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o contraseña invalidos ", 400)
		return
	}

	_, err = middlew.AuthLDAP(t.Email, t.Password)
	if err != nil {
		http.Error(w, "Autenticacion fallida con LDAP: "+err.Error(), 400)
		return
	}
	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el token: "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
