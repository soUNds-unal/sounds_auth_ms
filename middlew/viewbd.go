package middlew

import (
	"net/http"

	"github.com/ccmorenov/microservicesounds/bd"
)

func ViewBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ViewConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
