package mongocon

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	DB *mongo.Client
}

func ConnectMongoDB() *MongoDB {
	uri := os.Getenv("mongoURI")
	client, err := mongo.NewClient(options.Client().SetDirect(true).ApplyURI(uri))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Can not connect to the db server")
		//panic(err)
	}

	//fmt.Println("connection ok")

	return &MongoDB{client}
}
