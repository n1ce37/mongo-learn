package model

import (
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/mongodb/mongo-go-driver/bson"
)

var templateC *mongo.Collection

type Template struct {
	ID           string `json:"id"`
	CreateTime   int64  `json:"create_time"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
	Data         bson.A `json:"data"`
}

func (*Template) Init(db *mongo.Database) {
	templateC = db.Collection("template")
}

func (t *Template) Insert() (*mongo.InsertOneResult, error) {
	t.CreateTime = time.Now().Unix()
	return templateC.InsertOne(nil, t)
}
