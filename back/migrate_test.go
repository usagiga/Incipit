package main

import (
	"github.com/usagiga/Incipit/back/lib/testutils"
	"testing"
)

func TestMigrate(t *testing.T) {
	// Initialize db
	db, dbFin := testutils.SetupTestDB(t)
	defer dbFin()

	// Migrate
	err := Migrate(db)
	if err != nil {
		t.Error(err)
	}

	// Validate set tables
	// TODO : Implement it
}
