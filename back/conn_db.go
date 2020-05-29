package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/usagiga/Incipit/back/lib/broker"
)

func ConnectToDB(userName, password, host string, port int) *gorm.DB {
	connAddr := fmt.Sprintf("%s:%s@tcp(%s:%d)/incipit?charset=utf8mb4&parseTime=true", userName, password, host, port)
	dbChan := make(chan *gorm.DB)
	dbInit := broker.Default()

	go dbInit.Open(dbChan, "mysql", connAddr)

	return <- dbChan
}
