package login

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth struct {
	//ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `bson:"password"`
}

type Authentication struct {
	client *mongo.Client
	ctx    context.Context
}

func NewAuthentication(ctx context.Context, client *mongo.Client) Authentication {
	return Authentication{
		client: client, ctx: ctx,
	}
}

func (a Authentication) CheckLogin(auth Auth) (bool, error) {
	collection := a.client.Database("example").Collection("auth")
	filter := bson.D{{"username", auth.Username}}
	var r Auth

	err := collection.FindOne(a.ctx, filter).Decode(&r)
	if err != nil {
		return false, err
	}

	if r.Password == auth.Password {
		return true, nil
	}

	return false, nil
}
