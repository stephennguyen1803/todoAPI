package models

//Todo struck Represents a Todo, we uses bson keyword to tell the mgo driver
//how to name the properties in mongodb document
type Todo struct {
	ID          int    `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string `bson:"title,omitempty" json:"title,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}
