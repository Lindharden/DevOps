package test

import (
	"DevOps/globals"
	"log"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	gormDb, err := gorm.Open(sqlite.Open("file:testsGorm?mode=memory&cache=shared"), &gorm.Config{})
	// set the global db connection

	if err != nil {
		log.Fatal("Could not connect to database")
	}

	globals.SetDatabase(gormDb)

	//code that executes before full test suite

	exitVal := m.Run()
	//code that executes after full test suite

	os.Exit(exitVal)
}
