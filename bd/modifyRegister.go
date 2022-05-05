package bd

import (
	"context"
	"time"

	"github.com/ccmorenov/microservicesounds/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyRegister(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("sounds")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})

	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}

	registro["fechaNacimiento"] = u.FechaNacimiento

	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}

	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}

	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}

	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
