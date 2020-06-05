package broker

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"log"
	"math"
	"os"
	"time"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[Broker] ", log.LstdFlags)
}

type DBBroker struct {
	config *Config
}

func (b *DBBroker) Open(dbChan chan *gorm.DB, dialect string, args ...interface{}) {
	maxRetry := b.config.MaxRetry

	for i := 1; i <= maxRetry; i++ {
		// Try to conn
		db, err := gorm.Open(dialect, args...)

		// If error occurred, retry after exponential time
		if err != nil {
			wait := time.Duration(math.Pow(2.0, float64(i)))

			logger.Println("Error occurred establishing connection to DB.\nRetrying...\nDetails: ", err)
			time.Sleep(wait * time.Second)
			continue
		}

		// Configure DB
		b.configureDB(db.DB(), b.config)

		// Return result
		dbChan <- db
		return
	}

	// If timed out, raise an error
	panic("Can't establish connection to DB.")
}

func (b *DBBroker) configureDB(db *sql.DB, config *Config) {
	openConn := config.MaxOpenConn
	idleConn := config.MaxIdleConn
	maxLifetime := config.ConnMaxLifetime

	db.SetMaxOpenConns(openConn)
	db.SetMaxIdleConns(idleConn)
	db.SetConnMaxLifetime(maxLifetime)
}

type Config struct {
	MaxRetry int

	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime time.Duration
}

func New(config *Config) *DBBroker {
	return &DBBroker{config: config}
}

func Default() *DBBroker {
	defaultConfig := &Config{
		MaxRetry: 10,

		MaxOpenConn:     10,
		MaxIdleConn:     100,
		ConnMaxLifetime: 1 * time.Minute,
	}

	return New(defaultConfig)
}
