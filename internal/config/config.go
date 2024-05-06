package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	GRPCPort        string
	HTTPHost        string
	HTTPPort        int
	WebSocketHost   string
	WebSocketPort   int
	BackendAPIURL   string
	SessionDuration int
	RedisAddr       string
	RedisPassword   string
	RedisDB         int
}

func LoadConfig(configFilePath string) *Config {
	viper.SetConfigFile(configFilePath)
	_ = viper.ReadInConfig()

	viper.SetDefault("HTTP_HOST", "localhost")
	viper.SetDefault("HTTP_PORT", 8090)
	viper.SetDefault("WEBSOCKET_HOST", "localhost")
	viper.SetDefault("WEBSOCKET_PORT", 8091)
	viper.SetDefault("BACKEND_API_URL", "http://localhost:9999/")
	viper.SetDefault("SESSION_DURATION", 300)
	viper.SetDefault("GRPC_PORT", "50051")
	viper.SetDefault("REDIS_ADDR", "localhost:6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)

	viper.AutomaticEnv()

	return &Config{
		GRPCPort:        viper.GetString("GRPC_PORT"),
		HTTPPort:        viper.GetInt("HTTP_SERVER_PORT"),
		WebSocketHost:   viper.GetString("WEBSOCKET_HOST"),
		WebSocketPort:   viper.GetInt("WEBSOCKET_PORT"),
		BackendAPIURL:   viper.GetString("BACKEND_API_URL"),
		SessionDuration: viper.GetInt("SESSION_DURATION"),
		RedisAddr:       viper.GetString("REDIS_ADDR"),
		RedisPassword:   viper.GetString("REDIS_PASSWORD"),
		RedisDB:         viper.GetInt("REDIS_DB"),
	}
}
