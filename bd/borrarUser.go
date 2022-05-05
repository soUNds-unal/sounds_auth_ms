package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BorroUser(UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("sounds")
	col := db.Collection("usuarios")

	objID, _ := primitive.ObjectIDFromHex(UserID)

	condicion := bson.M{
		"_id": objID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}
