package data

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewDB(t *testing.T) {
	NewDB(nil)
}

func TestNewMongoDB(t *testing.T) {
	mongodb := NewMongoDB(nil)
	collection := mongodb.Database("kratos").Collection("test")
	saving0 := mongodbTest{Username: 123}
	saving1 := mongodbTest{Username: 1234444}
	saving2 := mongodbTest{Username: 8888}
	saving3 := mongodbTest{Username: 77777}
	some := []interface{}{saving0, saving1, saving2, saving3}
	many, err := collection.InsertMany(context.TODO(), some)
	if err != nil {
		spew.Dump("error: ", err)
	}
	spew.Dump("Log: ", many.InsertedIDs)

	spew.Dump("-------------------")

	var mongodbStruct mongodbTest
	hex, _ := primitive.ObjectIDFromHex("6333f873984a19c70e51ffe7")
	filter := bson.D{{"_id", hex}}
	if err := collection.FindOne(context.TODO(), filter).Decode(&mongodbStruct); err != nil {
		spew.Dump("error: ", err)
	}
	spew.Dump(mongodbStruct)
	assert.New(t).Equal(int64(77777), mongodbStruct.Username)
}

type mongodbTest struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Username int64
}
