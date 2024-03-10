package models

import (
	"go-todo/config"
	"time"
)

type Todos struct {
	Id          uint
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// function Create Todo for Post Method
func Create(todos Todos) bool {
	result, err := config.DB.Exec(`
		INSERT INTO Todos (Title, Description, CreatedAt, UpdatedAT)
		VALUE (?, ?, ?, ?)`, todos.Title, todos.Description, todos.CreatedAt, todos.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}
	lastInsert, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastInsert > 0
}

// we want to return just data
func GetAll() []Todos {
	rows, err := config.DB.Query(`SELECT * FROM Todos`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var todos []Todos

	for rows.Next() {
		var todo Todos
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			panic(err)
		}

		todos = append(todos, todo)
	}
	return todos
}
