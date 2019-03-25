package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"github.com/spf13/viper"
	"log"
)

type DBer interface {
	GetAll(collection string, selector interface{}, sort string, v interface{}) error
	GetOne(ctx context.Context, collection string, selector interface{}, v interface{}) error
	Insert(collection string, v interface{}) error
	Upsert(collection string, selector interface{}, v interface{}) error
}

func NewMongoDB() DBer {
	var err error
	var url string
	database := viper.GetString("db.mongodb.database")
	protocol := viper.GetString("db.mongodb.protocol")
	host := viper.GetString("db.mongodb.host")
	port := viper.GetString("db.mongodb.port")
	username := viper.GetString("db.mongodb.username")
	password := viper.GetString("db.mongodb.password")
	to := viper.GetDuration("db.mongodb.timeout")

	if username != "" && password != "" {
		url = fmt.Sprintf("%s%s:%s@%s:%s", protocol, username, password, host, port)
	} else {
		url = fmt.Sprintf("%s%s:%s", protocol, host, port)
	}

	ctx, _ := context.WithTimeout(context.Background(), to*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ = context.WithTimeout(context.Background(), 3*time.Second)
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &MongoDB{
		db: client.Database(database),
	}
}

type MongoDB struct {
	db *mongo.Database
}

func (m *MongoDB) GetAll(collection string, selector interface{}, sort string, v interface{}) error {
	return nil
}

func (m *MongoDB) GetOne(ctx context.Context, collection string, selector interface{}, v interface{}) error {
	c := m.db.Collection(collection)

	err := c.FindOne(ctx, selector).Decode(v)
	if err != nil {
		return err
	}

	return err
}

func (m *MongoDB) Insert(collection string, v interface{}) error {
	return nil
}

func (m *MongoDB) Upsert(collection string, selector interface{}, v interface{}) error {
	return nil
}
