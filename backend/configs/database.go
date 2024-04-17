package configs

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type DatabaseEnvConfig struct {
	DBHost          string
	DBUser          string
	DBName          string
	DBPort          string
	DBPassword      string
	DBConnectionUrl string
	DBType          string
	MustSeed        string
}

// ConnectToDatabase Connects to a database
func ConnectToDatabase(cfg *DatabaseEnvConfig, dbChan chan *gorm.DB, errChan chan error) {
	var err error
	dbErrChan := make(chan error, 1)
	go connectToDatabaseClient(cfg, dbErrChan)

	if err = <-dbErrChan; err != nil {
		dbChan <- nil
		errChan <- err
		return
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s  sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)

	if len(cfg.DBConnectionUrl) != 0 {
		dsn = cfg.DBConnectionUrl
	}

	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Info),
	//})

	log.Println("DSN2: ", dsn)

	sqlDB, err := sql.Open("pgx", dsn)

	if err != nil {
		dbChan <- nil
		errChan <- err
		return
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		dbChan <- nil
		errChan <- err
		return
	}

	dbChan <- db
	errChan <- nil

}

func connectToDatabaseClient(cfg *DatabaseEnvConfig, errChan chan error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s  sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword)

	if len(cfg.DBConnectionUrl) != 0 {
		dsn = cfg.DBConnectionUrl
	}

	log.Println("DSN: ", dsn)
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Printf("connectToDatabaseClient Err: %v", err)
		errChan <- err
		return
	}

	db.Exec(fmt.Sprintf("CREATE DATABASE %s", cfg.DBName))
	//execute query
	errChan <- nil

}
