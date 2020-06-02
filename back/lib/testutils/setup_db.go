package testutils

import (
	"github.com/jinzhu/gorm"
	"testing"
)

// SetupTestDB connects and migrates DB,
// then passes `Finalizer` which drops tables and closes connection
func SetupTestDB(t *testing.T) (db *gorm.DB, finalizer Finalizer) {
	// Initialize DB
	db, dbFin := ConnectToTestDB()
	err := MigrateTestDB(db)
	if err != nil {
		dbFin()
		t.Fatal("SetUpDB(): Can't migrate DB; ", err)
		return nil, nil
	}

	// Return result
	finalizer = func() {
		defer dbFin()

		// Drop tables
		err := DropTestDB(db)
		if err != nil {
			t.Fatal("SetUpDB(): Can't drop tables; ", err)
		}
	}

	return db, finalizer
}
