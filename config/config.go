package config

import (
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var config *viper.Viper

func InitializeConfig(logger *zap.SugaredLogger) {
	config = viper.New()

	config.SetConfigName("vesperis")
	config.SetConfigType("yml")
	config.AddConfigPath("./config")

	if _, err := os.Stat("./config/vesperis.yml"); os.IsNotExist(err) {
		logger.Warn("Config file not found, creating default config file.")
		createDefaultConfig()
	}

	if err := config.ReadInConfig(); err != nil {
		logger.Fatal("Error reading config file. - ", err)
	}

	logger.Info("Successfully created the config.")
}

func GetMySQLUrl() string {
	return config.GetString("databases.mysql.username") +
		":" +
		config.GetString("databases.mysql.password") +
		"@(" +
		config.GetString("databases.mysql.host") +
		":" +
		config.GetString("databases.mysql.port") +
		")/" +
		config.GetString("databases.mysql.database") +
		"?parseTime=true"
}

func GetRedisUrl() string {
	host := config.GetString("databases.redis.host")
	port := config.GetString("databases.redis.port")
	database := config.GetString("databases.redis.database")
	username := config.GetString("databases.redis.username")
	password := config.GetString("databases.redis.password")

	if username != "" && password != "" {
		return "redis://" + username + ":" + password + "@" + host + ":" + port + "/" + database
	}

	return "redis://" + host + ":" + port + "/" + database
}

func GetConfig() *viper.Viper {
	return config
}

func createDefaultConfig() {
	defaultConfig := []byte(`
databases:
  mysql:
    username: root
    password: password
    host: localhost
    port: 3306
    database: vesperis
  redis:
    host: localhost
    port: 6379
    database: 0
    username: ""
    password: ""
`)

	if err := os.MkdirAll("./config", os.ModePerm); err != nil {
		panic("Failed to create config directory: " + err.Error())
	}

	if err := os.WriteFile("./config/vesperis.yml", defaultConfig, 0644); err != nil {
		panic("Failed to create default config file: " + err.Error())
	}
}
