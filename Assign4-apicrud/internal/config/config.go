package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort     string
	MongoURI    string
	MongoDBName string
}

func New() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// set enviroment
	mongoHost := viper.GetString("MONGO_HOST")
	mongoPort := viper.GetString("MONGO_PORT")
	mongodbName := viper.GetString("MONGO_DATABASE")
	mongoUsername := viper.GetString("MONGO_USERNAME")
	mongoPassword := viper.GetString("MONGO_PASSWORD")
	sslmode := viper.GetString("MONGO_SSLMODE")

	if sslmode == "" {
		sslmode = "false"
	}

	if mongoUsername == "" || mongoPassword == "" {
		return nil, fmt.Errorf("MongoDB username or password is missing")
	}

	mongouri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
		mongoUsername, mongoPassword, mongoHost, mongoPort, mongodbName)

	// Return Config struct
	return &Config{
		AppPort:     viper.GetString("APP_PORT"),
		MongoURI:    mongouri,
		MongoDBName: mongodbName,
	}, nil
}
