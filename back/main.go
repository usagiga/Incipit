package main

import "log"

func main() {
	// Load config
	// TODO : Load from JSON
	mysqlUser := "incipit"
	mysqlPass := "incipit-password"
	mysqlHost := "db"
	mysqlPort := 3304

	// Connect to DB
	db := ConnectToDB(mysqlUser, mysqlPass, mysqlHost, mysqlPort)
	defer db.Close()

	// Auto migrate
	err := Migrate(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Build modules


	// Register to gin


	// Launch

}
