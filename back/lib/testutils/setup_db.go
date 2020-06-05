package testutils

import (
	"github.com/jinzhu/gorm"
	"testing"
)

// SetupTestDB connects and migrates DB,
// then passes `Finalizer` which drops tables and closes connection
func SetupTestDB(t *testing.T) (db *gorm.DB, finalizer Finalizer) {
	// Load config
	config, err := LoadTestConfig()
	if err != nil {
		t.Fatal("SetUpDB(): Can't load config; ", err)
		return nil, nil
	}

	// Initialize DB
	db, dbFin := ConnectToTestDB(config)
	err = MigrateTestDB(db)
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
