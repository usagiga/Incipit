package main

import (
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/lib/config"
	"log"
)

func main() {
	// Load config
	c := &entity.Config{}
	err := config.Load(c)
	if err != nil {
		log.Fatalln(err)
	}

	// Connect to DB
	db := ConnectToDB(c.MySQLUser, c.MySQLPassword, c.MySQLHost, c.MySQLPort)
	defer db.Close()

	// Auto migrate
	err = Migrate(db)
	if err != nil {
		log.Fatalln(err)
	}

	// Build modules

	// Register to gin

	// Launch

}
