// database/database.go
package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "golangDb"
const connectionString = "mongodb+srv://ritik:bY1MUy916TmqCu4o@golangdb.gqx5mtp.mongodb.net/?retryWrites=true&w=majority"

var UserCollection *mongo.Collection
var ProductCollection *mongo.Collection

func Initialize() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	UserCollection = client.Database(dbName).Collection("users")
	ProductCollection = client.Database(dbName).Collection("products")
}
