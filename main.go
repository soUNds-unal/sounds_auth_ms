package main

import (
	"log"

	"github.com/ccmorenov/microservicesounds/bd"
	"github.com/ccmorenov/microservicesounds/handlers"
)

func main() {
	if bd.ViewConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
