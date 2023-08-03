package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

)

type Netflix struct{
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`,
	Title string `json:"title,omitempty" bson:"title,omitempty"`,
	Year string `json:"year,omitempty" bson:"year,omitempty"`,
	Rating string `json:"rating,omitempty" bson:"rating,omitempty"`,
	Watched bool `json:"watched,omitempty" bson:"watched,omitempty"`,
}