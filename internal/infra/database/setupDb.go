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

var (
	mongoOnce     sync.Once
	MongoInstance *MongoDBConnection
)

func SetupMongoDB(config MongoDBConfig) *MongoDBConnection {
	mongoOnce.Do(func() {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.URI))
		if err != nil {
			log.Fatal("Error connecting to MongoDB: ", err)
		}

		MongoInstance = &MongoDBConnection{
			Client: client,
			Db:     client.Database(config.Database),
		}

		err = MongoInstance.Client.Ping(context.TODO(), options.Client().ReadPreference)
		if err != nil {
			log.Fatal("Ping fail! Can't connect to MongoDB. Check your credentials and system status.")
		}

		fmt.Println("MongoDB successfully connected.")
	})

	return MongoInstance
}
