package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*Objeto de conexion BD*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0.ofge1.mongodb.net/sounds?retryWrites=true&w=majority")

/*Funcion que me permite conectar la BD */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la BD.")
	return client
}

/*Ping BD*/
func ViewConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
