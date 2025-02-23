package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

// Connect with db
func ConnectDB() {
	mongoURI := AppConfig.MONGODB_URI
	if mongoURI == "" {
		log.Fatal("MONGODB_URI is required but not set in env")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error: ", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("MongoDB ping failed: ", err)
	}

	dbName := AppConfig.DB_NAME
	if dbName == "" {
		log.Fatal("DB_NAME is required but not set in env")
	}

	DB = client.Database(dbName)
	// fmt.Println("Connected to MongoDB...")
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Collection(collectionName)
}
