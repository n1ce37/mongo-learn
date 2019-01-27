package model

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

var templateC *mongo.Collection

type Template struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	CreateTime   int64              `json:"create_time,omitempty" bson:"create_time"`
	Title        string             `json:"title,omitempty" bson:"title"`
	Introduction string             `json:"introduction,omitempty" bson:"introduction"`
	Data         bson.A             `json:"data,omitempty" bson:"data"`
}

func (*Template) Init(db *mongo.Database) {
	templateC = db.Collection("template")
}

func (t *Template) Insert() (*mongo.InsertOneResult, error) {
	t.CreateTime = time.Now().Unix()
	t.ID = primitive.NewObjectID()
	return templateC.InsertOne(nil, t)
}

func (t *Template) FindMany(page, size int) ([]Template, int, error) {
	offset := getOffset(page, size)
	filter := bson.D{{}}
	sort := bson.D{
		{"create_time", -1},
	}

	cur, err := templateC.Find(
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
	if err != nil {
		return nil, -1, err
	}
	defer cur.Close(nil)

	cnt, err := templateC.Count(nil, filter)
	if err != nil {
		return nil, -1, err
	}

	templates := make([]Template, 0)
	for cur.Next(nil) {
		t := Template{}
		err = cur.Decode(&t)
		if err != nil {
			return nil, -1, err
		}
		templates = append(templates, t)
	}

	return templates, int(cnt), nil
}

func (t *Template) FindOne() error {
	filter := bson.D{
		{"_id", t.ID},
	}
	res := templateC.FindOne(nil, filter)
	err := res.Err()
	if err != nil {
		return err
	}
	err = res.Decode(t)
	if err != nil {
		return err
	}

	return nil
}
