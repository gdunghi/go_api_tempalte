package login

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Auth struct {
	//ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `bson:"password"`
}

type Authentication struct {
	db DBer
}

func NewAuthentication(db DBer) Authentication {
	return Authentication{
		db: db,
	}
}

func (a Authentication) CheckLogin(ctx context.Context, auth Auth) (bool, error) {
	filter := bson.D{{"username", auth.Username}}
	var r Auth

	err := a.db.GetOne(ctx, "auth", filter, &r)

	if err != nil {
		return false, err
	}

	if r.Password == auth.Password {
		return true, nil
	}

	return false, nil
}
