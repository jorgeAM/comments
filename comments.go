package main

import (
	"flag"
	"log"

	"github.com/jorgeAM/comments/migration"
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
}
