package test

import (
	"DevOps/globals"
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/tanimutomo/sqlfile"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	db, _ := sqlx.Open("sqlite3", "file:tests?mode=memory&cache=shared")
	gormDb, err := gorm.Open(sqlite.Open("file:testsGorm?mode=memory&cache=shared"), &gorm.Config{})
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
