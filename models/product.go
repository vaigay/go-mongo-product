package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Price float32            `json:"price" bson:"price"`
}
