package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/jorgeAM/comments/migration"
	"github.com/jorgeAM/comments/routes"
	"github.com/urfave/negroni"
)

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Generar las migraciones")
	flag.Parse()
	if migrate == "yes" {
		log.Println("comenzó la migración")
		migration.Migrate()
		log.Println("Finalizó la migración")
	}

	// Inicia las rutas
	router := routes.InitRoutes()

	// Inicia los middlewares
	n := negroni.Classic()
	n.UseHandler(router)

	// Inicia el Servidor
	server := &http.Server{
		Addr:    ":8000",
		Handler: n,
	}

	log.Println("Servidor corriendo en http://localhost:8000")
	log.Println(server.ListenAndServe())
}
