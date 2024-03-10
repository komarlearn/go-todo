package controllers

import (
	"go-todo/models"
	"net/http"
	"text/template"
	"time"
)

func Add(w http.ResponseWriter, r *http.Request) {
	// method GET untuk menampilkan
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/todos/create.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	}

	// method POST untuk menambahkan dan megirim
	if r.Method == "POST" {
		var todos models.Todos

		todos.Title = r.FormValue("title")
		todos.Description = r.FormValue("description")
		todos.CreatedAt = time.Now()
		todos.UpdatedAt = time.Now()

		if ok := models.Create(todos); !ok {
			temp, err := template.ParseFiles("views/todos/create.html")
			if err != nil {
				panic(err)
			}
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)

	}
}

// function return data
func Data(w http.ResponseWriter, r *http.Request) {
	todos := models.GetAll()
	data := map[string]any{
		"todos": todos,
	}
	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}
