package main

import (
	"github.com/gorilla/mux"
	"go-contacts/app"
	"os"
	"fmt"
	"net/http"
	"go-contacts/controllers"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JWTAuthentication) //attach JWT auth middleware

	router.HandleFunc("/api/user/new",
	controllers.CreateAccount).Methods("POST")

	router.HandleFunc("/api/user/login",
	controllers.Authenticate).Methods("POST")




	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" 
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000
	if err != nil {
		fmt.Print(err)
	}
}