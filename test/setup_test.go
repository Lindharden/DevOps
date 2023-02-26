package test

import (
	"DevOps/globals"
	"log"
	"os"
	"testing"

	"database/sql"

	"github.com/tanimutomo/sqlfile"
)

func TestMain(m *testing.M) {
	db, err := sql.Open("sqlite3", globals.DATABASE)

	if err != nil {
		log.Fatal("Could not connect to database")
	}

	s := sqlfile.New()

	// Load input file and store queries written in the file
	loaderr := s.File("../tools/schema.sql")
	if loaderr != nil {
		log.Fatal("Could not load database file")
	}
	_, err = s.Exec(db)

	if err != nil {
		log.Fatal("Error executing startup script")
	}

	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}
