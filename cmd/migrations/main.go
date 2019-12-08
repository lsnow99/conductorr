package main

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

const directory = "cmd/migrations"

func main() {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "conductorr",
		Database: "conductorrdb",
	})

	err := migrations.Run(db, directory, os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}