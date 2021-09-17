package mongocon

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	DB *mongo.Client
}

func ConnectMongoDB() *MongoDB {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return &MongoDB{client}
}
