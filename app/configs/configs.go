package configs

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Name string
	Env  string
	Port string
}

type DbConfigOne struct {
	Host        string
	Port        string
	Dbname      string
	Username    string
	Password    string
	DbIsMigrate bool
	DebugMode   bool
}
type DbConfigTwo struct {
	Host        string
	Port        string
	Dbname      string
	Username    string
	Password    string
	DbIsMigrate bool
	DebugMode   bool
}

type Configs struct {
	Appconfig   AppConfig
	Dbconfigone DbConfigOne
	Dbconfigtwo DbConfigTwo
}

var lock = &sync.Mutex{}
var configs *Configs

func GetInstance() *Configs {
	if configs == nil {
		lock.Lock()

		if err := godotenv.Load(); err != nil {
			log.Println("Failed to load env file")
		}

		configs = &Configs{
			Appconfig: AppConfig{
				Name: getEnv("APP_NAME", "challenge-lion"),
				Env:  getEnv("APP_ENV", "dev"),
				Port: getEnv("APP_PORT", "8000"),
			},

			Dbconfigone: DbConfigOne{
				Host:        getEnv("DB_HOST_ONE", "localhost"),
				Port:        getEnv("DB_PORT_ONE", "5432"),
				Username:    getEnv("DB_USER_ONE", "postgres"),
				Password:    getEnv("DB_PASSWORD_ONE", ""),
				Dbname:      getEnv("DB_NAME_ONE", "db_source"),
				DbIsMigrate: getEnv("DB_ISMIGRATE_ONE", "true") == "true",
				DebugMode:   getEnv("DEBUG_MODE_ONE", "false") == "true",
			},
			Dbconfigtwo: DbConfigTwo{
				Host:        getEnv("DB_HOST_TWO", "localhost"),
				Port:        getEnv("DB_PORT_TWO", "5432"),
				Username:    getEnv("DB_USER_TWO", "postgres"),
				Password:    getEnv("DB_PASSWORD_TWO", ""),
				Dbname:      getEnv("DB_NAME_TWO", "db_destination"),
				DbIsMigrate: getEnv("DB_ISMIGRATE_TWO", "true") == "true",
				DebugMode:   getEnv("DEBUG_MODE_TWO", "false") == "true",
			},
		}
		lock.Unlock()
	}

	return configs
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
