package routes

import (
	"github.com/JustinHaTran/ImageRepo/handlers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/user/", handlers.GetUser).Methods("GET")
	router.HandleFunc("/user/{userId}", handlers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", handlers.UpdateUser).Methods("PUT")
}