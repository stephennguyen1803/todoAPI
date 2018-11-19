package da

import (
	"log"

	. "todoAPI/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ToDoDataAccess struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "todo"
)

// Establish a connection to database
func (td *ToDoDataAccess) Connect() {
	session, err := mgo.Dial(td.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(td.Database)
}

// FindAll use to Find list of todolist
func (td *ToDoDataAccess) FindAll() ([]Todo, error) {
	var todolist []Todo
	err := db.C(COLLECTION).Find(bson.M{}).All(&todolist)
	return todolist, err
}

// Insert a task into database
func (td *ToDoDataAccess) Insert(todo Todo) error {
	err := db.C(COLLECTION).Insert(&todo)
	return err
}

// Delete use to remove task on todo list
func (td *ToDoDataAccess) Delete(todo Todo) error {
	err := db.C(COLLECTION).Remove(&todo)
	return err
}
