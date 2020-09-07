package testutils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/lib/broker"
)

func ConnectToTestDB(config *entity.Config) (db *gorm.DB, finalizer Finalizer) {
	connAddr := config.GetDSN()
	dbChan := make(chan *gorm.DB)
	dbInit := broker.Default()

	go dbInit.Open(dbChan, "mysql", connAddr)
	db = <-dbChan

	return db, func() {
		db.Close()
	}
}
