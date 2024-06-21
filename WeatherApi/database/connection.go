package database

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance      *mongo.Client
	clientInstanceError error
	mongoOnce           sync.Once
)

// mongodb://admin:password@mongodb

// mongodb://localhost:27017
const (
	CONNECTIONSTRING = "mongodb://admin:password@mongodb"
	DB               = "testdb"
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		clientInstance, clientInstanceError = mongo.Connect(context.TODO(), clientOptions)
		if clientInstanceError != nil {
			log.Fatal(clientInstanceError)
		}
		err := clientInstance.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}
	})
	return clientInstance, clientInstanceError
}

func GetCollection(collectionName string) (*mongo.Collection, error) {
	client, err := GetMongoClient()
	if err != nil {
		return nil, err
	}
	collection := client.Database(DB).Collection(collectionName)
	return collection, nil
}
