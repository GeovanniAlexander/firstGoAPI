package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/GeovanniAlexander/01-firstGoAPI/src/helpers"
	"github.com/GeovanniAlexander/01-firstGoAPI/src/models"
	"github.com/gorilla/mux"
)

type Data struct {
	Success bool          `json: Success`
	Data    []models.Todo `json:data`
	Errors  []string      `json:errors`
}

func CreateTodo(res http.ResponseWriter, req *http.Request) {
	bodyTodo, success := helpers.DecodeBody(req)
	if !success {
		http.Error(res, "couldnt decode body", http.StatusBadRequest)
		return
	}
	var data Data = Data{Errors: make([]string, 0)}
	bodyTodo.Description = strings.TrimSpace(bodyTodo.Description)
	if !helpers.IsValidDescription(bodyTodo.Description) {
		data.Success = false
		data.Errors = append(data.Errors, "Invalid Description")

		json, _ := json.Marshal(data)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(json)
		return
	}
	todo, success := models.Insert(bodyTodo.Description)
	if !success {
		data.Errors = append(data.Errors, "Couldnt create todo")
	}
	data.Success = true
	data.Data = append(data.Data, todo)
	json, _ := json.Marshal(data)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(json)
	return
}

func GetTodo(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var data Data
	var todo, success = models.Get(id)
	if !success {
		data.Success = false
		data.Errors = append(data.Errors, "Todo not found")

		json, _ := json.Marshal(data)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(json)
		return
	}
	data.Success = true
	data.Data = append(data.Data, todo)
	json, _ := json.Marshal(data)
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(json)

	return
}

func GetTodos(res http.ResponseWriter, req *http.Request) {
	var todos []models.Todo = models.GetAll()
	var data = Data{true, todos, make([]string, 0)}
	json, _ := json.Marshal(data)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(json)

	return
}

func DeleteTodo(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	var data = Data{Errors: make([]string, 0)}
	todo, success := models.Delete(id)
	if !success {
		data.Errors = append(data.Errors, "Couldn delete todo")
	}
	data.Success = true
	data.Data = append(data.Data, todo)
	json, _ := json.Marshal(data)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(json)

	return
}

func UpdateTodo(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	bodyTodo, success := helpers.DecodeBody(req)
	if !success {
		http.Error(res, "couldnt decode body", http.StatusBadRequest)
		return
	}
	var data Data = Data{Errors: make([]string, 0)}
	bodyTodo.Description = strings.TrimSpace(bodyTodo.Description)
	if !helpers.IsValidDescription(bodyTodo.Description) {
		data.Success = false
		data.Errors = append(data.Errors, "Invalid description")
		json, _ := json.Marshal(data)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(json)
		return
	}
	todo, success := models.Update(id, bodyTodo.Description)
	if !success {
		data.Errors = append(data.Errors, "Couldnt update todos")
	}
	data.Success = true
	data.Data = append(data.Data, todo)
	json, _ := json.Marshal(data)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(json)

	return
}
