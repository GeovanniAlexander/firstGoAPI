package models

import (
	"github.com/GeovanniAlexander/01-firtsGoAPI/src/database"
)

type Todo struct {
	ID          int    `json:ID`
	Description string `json:description`
}

func Insert(description string) (Todo, bool) {
	db := database.GetConnection()
	var todo_id int
	db.QueryRow("INSERT INTO todos(description) VALUES($1) RETURNING id", description).Scan(&todo_id)

	if todo_id == 0 {
		return Todo{}, false
	}
	return Todo{todo_id, description}, true
}

func Get(id string) (Todo, bool) {
	db := database.GetConnection()
	row := db.QueryRow("SELECT * FROM todos WHERE id = $1", id)

	var ID int
	var description string
	err := row.Scan(&ID, &description)
	if err != nil {
		return Todo{}, false
	}

	return Todo{ID, description}, true
}
