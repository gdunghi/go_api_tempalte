package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/spf13/viper"
	"log"
)

type MongoDB struct {
	Client *mongo.Client
	dbName string
}

func NewMongoDB(ctx context.Context) *MongoDB {
	var err error
	var url string
	database := viper.GetString("db.mongodb.database")
	protocol := viper.GetString("db.mongodb.protocol")
	host := viper.GetString("db.mongodb.host")
	port := viper.GetString("db.mongodb.port")
	username := viper.GetString("db.mongodb.username")
	password := viper.GetString("db.mongodb.password")

	if username != "" && password != "" {
		url = fmt.Sprintf("%s%s:%s@%s:%s", protocol, username, password, host, port)
	} else {
		url = fmt.Sprintf("%s%s:%s", protocol, host, port)
	}
	log.Println(url)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &MongoDB{
		Client: client,
		dbName: database,
	}
}
