package da

import (
	"log"
	"todoAPI/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ToDoDataAccess is struck to connect database
type ToDoDataAccess struct {
	Host     string
	Database string
}

var db *mgo.Database

// name collection
const (
	COLLECTION = "todolist"
)

// Connect create a connect
func (td *ToDoDataAccess) Connect() {
	session, err := mgo.Dial(td.Host)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(td.Database)
}

// FindAll use to Find list of todolist
func (td *ToDoDataAccess) FindAll() ([]models.Todo, error) {
	var todolist []models.Todo
	err := db.C(COLLECTION).Find(bson.M{}).All(&todolist)
	return todolist, err
}

// Insert a task into database
func (td *ToDoDataAccess) Insert(todo models.Todo) error {
	err := db.C(COLLECTION).Insert(todo)
	return err
}

// Delete use to remove task on todo list
func (td *ToDoDataAccess) Delete(todo models.Todo) error {
	err := db.C(COLLECTION).RemoveId(todo.ID)
	return err
}

// Update a task into database
func (td *ToDoDataAccess) Update(todo models.Todo) error {
	err := db.C(COLLECTION).UpdateId(todo.ID, bson.M{
		"$set": todo,
	})
	return err
}
