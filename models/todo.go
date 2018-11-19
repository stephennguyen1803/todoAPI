package models

import "gopkg.in/mgo.v2/bson"

//Todo struck Represents a Todo, we uses bson keyword to tell the mgo driver
//how to name the properties in mongodb document
type Todo struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
}
