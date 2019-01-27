package model

import (
	"context"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var historyC *mongo.Collection

type History struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	UUID         string             `json:"uuid,omitempty" bson:"uuid"`
	CreateTime   int64              `json:"create_time,omitempty" bson:"create_time"`
	Title        string             `json:"title,omitempty" bson:"title"`
	Introduction string             `json:"introduction,omitempty" bson:"introduction"`
	Data         bson.A             `json:"data,omitempty" bson:"data"`
}

func (*History) Init(db *mongo.Database) {
	historyC = db.Collection("history")
}

func (h *History) Insert() (*mongo.InsertOneResult, error) {
	h.CreateTime = time.Now().Unix()
	return historyC.InsertOne(nil, h)
}

func (h *History) FindMany(offset, limit int) ([]History, int, error) {
	cur, err := historyC.Find(
		nil,
		bson.D{
			{
				"uuid", h.UUID,
			},
		},
		options.Find().SetProjection(bson.D{
			{"_id", 1},
			{"title", 1},
			{"introduction", 1},
		}),
		options.Find().SetSkip(int64(offset)),
		options.Find().SetLimit(int64(limit)),
	)

	cnt, err := historyC.Count(
		nil,
		bson.D{
			{
				"uuid", h.UUID,
			},
		},
	)

	if err != nil {
		return nil, -1, err
	}

	defer cur.Close(context.Background())

	histories := make([]History, 10)
	for i := 0; i < limit; i++ {
		if cur.Next(nil) {
			err := cur.Decode(&histories[i])
			if err != nil {
				log.Println(err)
				break
			}
		}
	}

	return histories, int(cnt), nil

}
