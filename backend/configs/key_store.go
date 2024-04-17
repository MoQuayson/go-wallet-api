package configs

import (
	"github.com/gobuffalo/envy"
	"go-wallet-api/features/shared/utils"
	"gorm.io/gorm"
	"log"
	"os"
)

var WithKeyStoreConfig *KeyStoreConfig

type KeyStoreConfig struct {
	DBConfig  *DatabaseEnvConfig
	JWTSecret string
	APIPort   string
	DBContext *gorm.DB
}

func NewKeyStoreConfig() *KeyStoreConfig {
	cfg := &KeyStoreConfig{}
	dbChan, errChan := utils.MakeDataAndErrorChannels[gorm.DB]()

	cfg.DBConfig = loadDBConfigFromEnv()
	cfg.JWTSecret = os.Getenv("JWT_SECRET")
	cfg.APIPort = envy.Get("API_PORT", "8080")
	go ConnectToDatabase(cfg.DBConfig, dbChan, errChan)

	if err := <-errChan; err != nil {
		log.Fatalln("failed to connect to database")
		return nil
	}

	cfg.DBContext = <-dbChan
	log.Println("connected to database")
	return cfg

}

func loadDBConfigFromEnv() *DatabaseEnvConfig {
	return &DatabaseEnvConfig{
		DBHost:          envy.Get("DB_HOST", "127.0.0.1"),
		DBUser:          envy.Get("DB_USER", "postgres"),
		DBName:          envy.Get("DB_NAME", "wallet_db"),
		DBPort:          envy.Get("DB_PORT", "5432"),
		DBPassword:      envy.Get("DB_PASSWORD", ""),
		DBConnectionUrl: envy.Get("DB_CONNECTION_URL", ""),
	}
}
