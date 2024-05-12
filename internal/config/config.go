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
	MongoIP                 string
	MongoPort               int
	MongoUser               string
	MongoPassword           string
	MongoAuthSource         string
	MongoDBName             string
	MongoActionCollection   string
	MatchServiceGRPCAddress string
}

func LoadConfig(configFilePath string) *Config {
	viper.SetConfigFile(configFilePath)
	if err := viper.ReadInConfig(); err != nil {
		log.Warn(err)
	}

	viper.SetDefault("GRPC_PORT", "50051")
	viper.SetDefault("HTTP_HOST", "localhost")
	viper.SetDefault("HTTP_PORT", 30002)
	viper.SetDefault("WEBSOCKET_HOST", "localhost")
	viper.SetDefault("WEBSOCKET_PORT", 30003)
	viper.SetDefault("SESSION_DURATION", 300)
	viper.SetDefault("REDIS_ADDR", "localhost:6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("MONGO_IP", "localhost")
	viper.SetDefault("MONGO_PORT", 27017)
	viper.SetDefault("MONGO_USER", "root")
	viper.SetDefault("MONGO_PASSWORD", "password")
	viper.SetDefault("MONGO_AUTH_SOURCE", "admin")
	viper.SetDefault("MONGO_DB_NAME", "assistant")
	viper.SetDefault("MONGO_ACTION_COLLECTION", "scenario")
	viper.SetDefault("MATCH_SERVICE_GRPC_ADDRESS", "localhost:50051")
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
		MongoIP:                 viper.GetString("MONGO_IP"),
		MongoPort:               viper.GetInt("MONGO_PORT"),
		MongoUser:               viper.GetString("MONGO_USER"),
		MongoPassword:           viper.GetString("MONGO_PASSWORD"),
		MongoAuthSource:         viper.GetString("MONGO_AUTH_SOURCE"),
		MongoDBName:             viper.GetString("MONGO_DB_NAME"),
		MongoActionCollection:   viper.GetString("MONGO_ACTION_COLLECTION"),
		MatchServiceGRPCAddress: viper.GetString("MATCH_SERVICE_GRPC_ADDRESS"),
	}
}
