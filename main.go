package main

import (
	"flag"
	"log"
	"net/http"
	das "todoAPI/da"
	"todoAPI/models"

	"github.com/facebookgo/flagenv"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var da = das.ToDoDataAccess{}
var mgoURI string
var databaseName string

func int() {

	godotenv.Load()

	flag.StringVar(&mgoURI, "mgo-uri", "localhost", "uri to connect to mongodb")
	flag.StringVar(&databaseName, "database-name", "todolist123", "database use to work")
	flag.Parse()

	flagenv.Parse()

	da.Host = mgoURI
	da.Database = databaseName
	da.Connect()
}

// our main function
func main() {
	router := gin.Default()

	router.GET("/", GetAll)
	router.POST("/createTodo", CreateTodoEndPoint)
	router.PUT("/updateTodo", UpdateTodoEndPoint)
	router.DELETE("/deleteTodo", DeleteTodoEndPost)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

// GetAll use to gell all Todo data
func GetAll(c *gin.Context) {
	todo, err := da.FindAll()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, todo)
}

// CreateTodoEndPoint use to create task
func CreateTodoEndPoint(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := da.Insert(todo); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, todo)

}

// UpdateTodoEndPoint use to update data
func UpdateTodoEndPoint(c *gin.Context) {

	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "err.Error()"})
		return
	}

	if err := da.Update(todo); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)

}

// DeleteTodoEndPost use to remove task by id
func DeleteTodoEndPost(c *gin.Context) {

	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := da.Delete(todo); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, "Remove success todo with id %s", todo.ID)
}
