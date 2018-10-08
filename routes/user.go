package routes

import (
	"github.com/gorilla/mux"
	"github.com/jorgeAM/comments/controllers"
	"github.com/urfave/negroni"
)

// SetUserRouter rutas de user
func SetUserRouter(router *mux.Router) {
	prefix := "/api/users"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controllers.UserCreate).Methods("POST")
	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
