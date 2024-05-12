package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	GRPCPort                string
	HTTPHost                string
	HTTPPort                int
	WebSocketHost           string
	WebSocketPort           int
	BackendAPIURL           string
	SessionDuration         int
	RedisAddr               string
	RedisPassword           string
	RedisDB                 int
	MongoURI                string
	MongoDBName             string
	MongoActionCollection   string
	MatchServiceGRPCAddress string
}

func LoadConfig(configFilePath string) *Config {
	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Warn(err)
	}

	viper.SetDefault("HTTP_HOST", "localhost")
	viper.SetDefault("HTTP_PORT", 8080)
	viper.SetDefault("WEBSOCKET_HOST", "localhost")
	viper.SetDefault("WEBSOCKET_PORT", 8081)
	viper.SetDefault("BACKEND_API_URL", "http://localhost:9999/")
	viper.SetDefault("SESSION_DURATION", 300)
	viper.SetDefault("GRPC_PORT", "50051")
	viper.SetDefault("REDIS_ADDR", "localhost:6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("MONGO_URI", "mongodb://localhost:27017")
	viper.SetDefault("MONGO_DB_NAME", "mydatabase")
	viper.SetDefault("MONGO_ACTION_COLLECTION", "actions")
	viper.SetDefault("MATCH_SERVICE_GRPC_ADDERSS", "localhost:50051")
	viper.AutomaticEnv()

	return &Config{
		GRPCPort:                viper.GetString("GRPC_PORT"),
		HTTPHost:                viper.GetString("HTTP_HOST"),
		HTTPPort:                viper.GetInt("HTTP_PORT"),
		WebSocketHost:           viper.GetString("WEBSOCKET_HOST"),
		WebSocketPort:           viper.GetInt("WEBSOCKET_PORT"),
		BackendAPIURL:           viper.GetString("BACKEND_API_URL"),
		SessionDuration:         viper.GetInt("SESSION_DURATION"),
		RedisAddr:               viper.GetString("REDIS_ADDR"),
		RedisPassword:           viper.GetString("REDIS_PASSWORD"),
		RedisDB:                 viper.GetInt("REDIS_DB"),
		MongoURI:                viper.GetString("MONGO_URI"),
		MongoDBName:             viper.GetString("MONGO_DB_NAME"),
		MongoActionCollection:   viper.GetString("MONGO_ACTION_COLLECTION"),
		MatchServiceGRPCAddress: viper.GetString("MATCH_SERVICE_GRPC_ADDERSS"),
	}
}
