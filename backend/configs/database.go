package configs

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseEnvConfig struct {
	DBHost     string
	DBUser     string
	DBName     string
	DBPort     string
	DBPassword string
	DBType     string
	MustSeed   string
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
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

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		errChan <- err
		return
	}

	db.Exec(fmt.Sprintf("CREATE DATABASE %s", cfg.DBName))
	//execute query
	errChan <- nil

}
