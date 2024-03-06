package config

import (
	"fmt"
	"go-wallet-api/database"
	"log"

	"github.com/gobuffalo/envy"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DbCtx *gorm.DB

type DatabaseConfig struct {
	Host     string
	User     string
	Name     string
	Port     string
	Password string
}

// Connects to a database
func ConnectDatabase() {
	var err error

	cfg := DatabaseConfig{
		Host:     envy.Get("DB_HOST", "127.0.0.1"),
		User:     envy.Get("DB_USER", "postgres"),
		Name:     envy.Get("DB_NAME", "wallet_db"),
		Port:     envy.Get("DB_PORT", "5432"),
		Password: envy.Get("DB_PASSWORD", ""),
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s  sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Name, cfg.Password)

	DbCtx, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln(err)
	}

	//migrate tables
	//err = database.MigrateEntities(DbCtx)
	//
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}

	//seed admin
	seedAdmin := envy.Get("SEED_ADMIN", "true")
	if seedAdmin == "true" {
		//Seed Admin
		err = database.SeedUserAdmin(DbCtx)

		if err != nil {
			log.Fatalln(err)
			return
		}
	}

}
