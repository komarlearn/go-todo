package controllers

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
