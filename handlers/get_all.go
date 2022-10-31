package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/weverton-souza/go-todo-task-crud/models"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Error: Internal Server Error:: Fail to retrieve data: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
