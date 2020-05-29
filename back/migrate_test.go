package main

import "testing"

func TestMigrate(t *testing.T) {
	mysqlUser := "incipit"
	mysqlPass := "incipit-password"
	mysqlHost := "db"
	mysqlPort := 3306

	// Connect to DB
	db := ConnectToDB(mysqlUser, mysqlPass, mysqlHost, mysqlPort)

	// Migrate
	err := Migrate(db)
	if err != nil {
		t.Error(err)
	}

	// Validate set tables
	// TODO : Implement it
}
