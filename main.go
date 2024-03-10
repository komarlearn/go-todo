package main

import (
	"fmt"
	"go-todo/config"
	"go-todo/controllers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Komar Ganteng, Todo aplication")

	// panggil connection ke database function
	config.ConnectDB()

	// static file
	http.Handle("/assets/",
		http.StripPrefix("/assets/",
			http.FileServer(http.Dir("assets"))))

	// route API or data
	http.HandleFunc("/", controllers.Data)

	// route views or UI/UX
	// http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/create", controllers.Add)

	// Log fo Listening port and connect to HTTP
	log.Println("Server Running on Port 8000")
	http.ListenAndServe(":8000", nil)
}
