package routes

import (
	"github.com/JustinHaTran/ImageRepo/handlers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{userId}", handlers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{userId}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/image/", handlers.CreateImage).Methods("POST")
	router.HandleFunc("/image/{imageId}", handlers.GetImageById).Methods("GET")
	router.HandleFunc("/image/{imageId}", handlers.UpdateImage).Methods("PUT")
}