package test

import (
	"DevOps/globals"
	"DevOps/helpers"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/tanimutomo/sqlfile"
)

func TestMain(m *testing.M) {
	db, err := sqlx.Open("sqlite3", "file:tests?mode=memory&cache=shared")
	gormDb := helpers.GetGormDbConnection()
	// set the global db connection

	globals.SetDatabase(db, gormDb)

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
	//code that executes before full test suite

	exitVal := m.Run()
	//code that executes after full test suite

	os.Exit(exitVal)
}
