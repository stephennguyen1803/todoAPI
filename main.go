package main

import (
	"encoding/json"
	"log"
	"net/http"

	. "todoAPI/da"
	. "todoAPI/models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var da = ToDoDataAccess{}

func init() {
	da.Server = "localhost"
	da.Database = "todo"
	da.Connect()
}

// our main function
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/todo", GetAll).Methods("GET")
	router.HandleFunc("/todo", CreateTodoEndPoint).Methods("POST")
	router.HandleFunc("/todo", DeleteTodoEndPost).Methods("DELETE")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}

}

// GetAll use to gell all Todo data
func GetAll(w http.ResponseWriter, r *http.Request) {
	todo, err := da.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, todo)
}

// CreateTodoEndPoint use to create task by id
func CreateTodoEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	todo.ID = bson.NewObjectId()
	if err := da.Insert(todo); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, todo)
}

// DeleteTodoEndPost use to remove task by id
func DeleteTodoEndPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := da.Delete(todo); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// respondWithJson return data type json
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
