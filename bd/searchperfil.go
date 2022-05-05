package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/ccmorenov/microservicesounds/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SearchPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("sounds")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return perfil, nil
	}
	return perfil, nil
}
