package model

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var historyC *mongo.Collection

type History struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
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
	h.ID = primitive.NewObjectID()
	return historyC.InsertOne(nil, h)
}

func (h *History) FindMany(page, size int) ([]History, int, error) {
	offset := getOffset(page, size)
	filter := bson.D{
		{
			"uuid", h.UUID,
		},
	}

	sort := bson.D{
		{"create_time", -1},
	}
	cur, err := historyC.Find(
		nil,
		filter,
		options.Find().SetProjection(bson.D{
			{"title", 1},
			{"introduction", 1},
		}),
		options.Find().SetSkip(int64(offset)),
		options.Find().SetLimit(int64(size)),
		options.Find().SetSort(sort),
	)

	cnt, err := historyC.Count(nil, filter)

	if err != nil {
		return nil, -1, err
	}

	defer cur.Close(context.Background())

	histories := make([]History, 0)
	for cur.Next(nil) {
		h := History{}
		err = cur.Decode(&h)
		if err != nil {
			return nil, -1, err
		}
		histories = append(histories, h)
	}

	return histories, int(cnt), nil
}

func (h *History) FindOne() error {
	filter := bson.D{
		{
			"_id", h.ID,
		},
	}
	opt := options.FindOne().SetProjection(
		bson.D{
			{"uuid", 0},
		},
	)
	res := historyC.FindOne(nil, filter, opt)
	err := res.Err()
	if err != nil {
		return err
	}
	err = res.Decode(h)
	if err != nil {
		return err
	}

	return nil
}

func (h *History) UpdateOne() error {
	filter := bson.D{
		{
			"_id", h.ID,
		},
	}
	updateData := bson.D{
		{"$set", bson.D{
			{"title", h.Title},
			{"introduction", h.Introduction},
			{"data", h.Data},
		}},
	}

	_, err := historyC.UpdateOne(nil, filter, updateData)
	if err != nil {
		return err
	}

	return nil
}

func (h *History) DeleteOne() error {
	filter := bson.D{
		{"_id", h.ID},
	}

	_, err := historyC.DeleteOne(nil, filter)
	return err
}
