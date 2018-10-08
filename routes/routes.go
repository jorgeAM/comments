package routes

import "github.com/gorilla/mux"

// InitRoutes inicializar routas
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	SetLoginRouter(router)
	SetUserRouter(router)
	return router
}
