package OperationsWithMongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

var (
	once     sync.Once
	mongoCli *mongo.Client
	err      error
)

func initMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	mongoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoCli.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(colorPurple + "Connected to MongoDB!" + colorReset)
	}
}

func GetMongoClient() *mongo.Client {
	once.Do(initMongoDB)
	return mongoCli
}

func GetDBName() string {
	return "irmag_innovations"
}

func GetCollectionName() string {
	return "products"
}

func DisconnectMongoDB(mongoCli *mongo.Client) {
	if mongoCli != nil {
		err := mongoCli.Disconnect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(colorPurple + "Disconnected from MongoDB" + colorReset)
	}
}
