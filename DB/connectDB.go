package DB

import (
	"context"

	"github.com/AnimeTrackerr/v2/backend/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(URI string) utils.ClientInfo {
	// define connection context like timeout and others if necessary
	context := context.TODO()

	client, err := mongo.Connect(context, options.Client().ApplyURI(URI))
	if err != nil {
		panic(err)
	}
	
	return utils.ClientInfo {
		Client: client,
		Ctx: context,
	}
}

func GetCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	return client.Database(dbName).Collection(collectionName)
}