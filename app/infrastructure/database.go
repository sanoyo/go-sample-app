package infrastructure

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

type Config struct {
	MySQL MySQLConfig
}

type MySQLConfig struct {
	Username string
	Password string
	Host     string
	Database string
}

func init() {
	var config Config

	var configFilePath string
	basePath := os.Getenv("VQR_PATH")

	env := os.Getenv("VQR_ENV")
	switch env {
	case "development":
		configFilePath = basePath + "config/config_dev.toml"
	default:
		configFilePath = basePath + "config/config_dev.toml"
	}

	_, err := toml.DecodeFile(configFilePath, &config)
	if err != nil {
		panic(err.Error())
	}

	dbConfig := config.MySQL.Username + ":" + config.MySQL.Password + "@tcp(" + config.MySQL.Host + ":3306)/" + config.MySQL.Database + "?parseTime=true"

	db, err := sqlx.Connect("mysql", dbConfig)
	if err != nil {
		panic(err.Error())
	}

	DB = db
}
