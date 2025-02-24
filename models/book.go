package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title    string             `bson:"title" json:"title"`
	Author   string             `bson:"author" json:"author"`
	Quantity int                `bson:"quantity" json:"quantity"`
}
