package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection = ConectarBD()
var clientsOptions = options.Client().ApplyURI("mongodb+srv://durbonca:elGueboAnalfabeta@cluster0.idou8.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* Connect to DB */
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientsOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión Exitosa a la base de datos")
	return client
}

func CheckConnection() int {
	err := MongoConnection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return 1
}
