package test

import (
	"DevOps/globals"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/tanimutomo/sqlfile"
)

func TestMain(m *testing.M) {
	db, err := sqlx.Open("sqlite3", "file:tests?mode=memory&cache=shared")

	// set the global db connection
	globals.DB = db

	if err != nil {
		log.Fatal("Could not connect to database")
	}

	s := sqlfile.New()

	// Load input file and store queries written in the file
	loaderr := s.File("../tools/schema.sql")
	if loaderr != nil {
		log.Fatal("Could not load database file")
	}
	_, err = s.Exec(db.DB)

	if err != nil {
		log.Fatal("Error executing startup script")
	}

	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}
