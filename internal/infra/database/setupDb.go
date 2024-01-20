package database

import (
	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConfig struct {
	URI      string
	Database string
}

type MongoDBConnection struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mongoOnce sync.Once
var mongoInstance *MongoDBConnection

func SetupMongoDB(config MongoDBConfig) *MongoDBConnection {
	mongoOnce.Do(func() {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.URI))
		if err != nil {
			log.Fatal("Error connecting to MongoDB: ", err)
		}

		mongoInstance = &MongoDBConnection{
			Client: client,
			Db:     client.Database(config.Database),
		}

		fmt.Println("MongoDB successfully connected.")
	})

	return mongoInstance
}
