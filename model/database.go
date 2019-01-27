package model

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var db *mongo.Database

func init() {
	client, err := mongo.NewClient("mongodb://root:example@localhost:27017")
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	db = client.Database("aws")

	initCollections()
}

func initCollections() {
	t := Template{}
	t.Init(db)
	h := History{}
	h.Init(db)
}
