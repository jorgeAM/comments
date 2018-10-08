package routes

import (
	"github.com/gorilla/mux"
	"github.com/jorgeAM/comments/controllers"
)

// SetLoginRouter rutas de user
func SetLoginRouter(router *mux.Router) {
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
}
