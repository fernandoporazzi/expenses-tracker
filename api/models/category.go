package models

import "gopkg.in/mgo.v2/bson"

// Category represents a document in the category collection
type Category struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	Name  string        `bson:"name" json:"name"`
	Slug  string        `bson:"slug" json:"slug"`
	Color string        `bson:"color" json:"color"`
}
