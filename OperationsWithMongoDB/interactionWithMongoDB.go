package OperationsWithMongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	colorReset  = "\033[0m"
	colorCyan   = "\033[36m"
	colorGreen  = "\033[32m"
	colorPurple = "\033[35m"
	colorRed    = "\033[31m"
)

func GetOrCreateCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	databases, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(colorRed, err, colorReset)
	}

	dbExists := false
	for _, db := range databases {
		if db == dbName {
			dbExists = true
			break
		}
	}

	db := client.Database(dbName)
	if !dbExists {
		db.Collection(collectionName)
	} else {
		collections, err := db.ListCollectionNames(context.TODO(), bson.M{})
		if err != nil {
			log.Fatal(colorRed, err, colorReset)
		}

		collectionExists := false
		for _, coll := range collections {
			if coll == collectionName {
				collectionExists = true
				break
			}
		}

		if !collectionExists {
			db.Collection(collectionName)
		}
	}

	return db.Collection(collectionName)
}

func InsertDataToMongoDB(ctx context.Context, item map[string]interface{}) {
	client := GetMongoClient()
	dbName := GetDBName()
	collectionName := GetCollectionName()

	collection := GetOrCreateCollection(client, dbName, collectionName)

	filter := bson.D{{"title", item["title"]}}

	var result bson.M
	errFind := collection.FindOne(context.TODO(), filter).Decode(&result)
	if errFind != nil {
		if errFind == mongo.ErrNoDocuments {
			if item != nil {
				_, err := collection.InsertOne(ctx, item)
				if err != nil {
					log.Println(colorRed, "Failed to insert item:", err, colorGreen, item["title"], colorReset)
				} else {
					log.Println(colorCyan, "Inserted item:", colorPurple, item["title"], colorReset)
				}
			}
		} else {
			log.Println(colorRed, errFind, colorReset)
		}
	} else {
		log.Println(colorRed, "There is already an item:", colorGreen, item["title"], colorReset)
	}
}

func GetDataByField(ctx context.Context, category string) chan map[string]interface{} {
	dataChannel := make(chan map[string]interface{})
	errorChannel := make(chan error, 1)

	go func() {
		defer close(dataChannel)
		defer close(errorChannel)

		client := GetMongoClient()
		collection := client.Database("irmag_innovations").Collection("products")

		cursor, err := collection.Find(ctx, bson.M{"category": category})
		if err != nil {
			errorChannel <- err
			return
		}
		defer cursor.Close(ctx)

		var phones []bson.M
		if err = cursor.All(ctx, &phones); err != nil {
			errorChannel <- err
			return
		}

		for _, value := range phones {
			if value["title"] != nil {
				select {
				case dataChannel <- value:
				case <-ctx.Done():
					errorChannel <- ctx.Err()
					return
				}
			}

		}
	}()

	return dataChannel
}
