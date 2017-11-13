package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Expense represent an expense in the expenses collection
type Expense struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	Date       time.Time     `bson:"date" json:"date"`
	Value      float32       `bson:"value" json:"value"`
	Categories []Category    `bson:"categories" json:"categories"`
}
