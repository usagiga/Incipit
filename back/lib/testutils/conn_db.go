package testutils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/lib/broker"
)

func ConnectToTestDB(config *entity.Config) (db *gorm.DB, finalizer Finalizer) {
	connAddr := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/incipit?charset=utf8mb4&parseTime=true",
		config.MySQLUser,
		config.MySQLPassword,
		config.MySQLHost,
		config.MySQLPort,
	)
	dbChan := make(chan *gorm.DB)
	dbInit := broker.Default()

	go dbInit.Open(dbChan, "mysql", connAddr)
	db = <-dbChan

	return db, func() {
		db.Close()
	}
}
