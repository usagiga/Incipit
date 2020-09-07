package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/usagiga/Incipit/back/entity"
	"github.com/usagiga/Incipit/back/lib/broker"
)

func ConnectToDB(config *entity.Config) *gorm.DB {
	connAddr := config.GetDSN()
	dbChan := make(chan *gorm.DB)
	dbInit := broker.Default()

	go dbInit.Open(dbChan, "mysql", connAddr)

	return <- dbChan
}
