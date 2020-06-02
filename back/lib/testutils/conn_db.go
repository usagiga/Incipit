package testutils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/usagiga/Incipit/back/lib/broker"
)

const (
	dbUserName = "incipit"
	dbPassword = "incipit-password"
	dbHost = "db"
	dbPort = 3306
)

func ConnectToTestDB() (db *gorm.DB, finalizer Finalizer) {
	connAddr := fmt.Sprintf("%s:%s@tcp(%s:%d)/incipit?charset=utf8mb4&parseTime=true", dbUserName, dbPassword, dbHost, dbPort)
	dbChan := make(chan *gorm.DB)
	dbInit := broker.Default()

	go dbInit.Open(dbChan, "mysql", connAddr)
	db = <-dbChan

	return db, func() {
		db.Close()
	}
}
