package main

import (
	"flag"

	"github.com/Korpenter/assistand/internal/app"
	"github.com/Korpenter/assistand/internal/config"
	"github.com/Korpenter/assistand/internal/hub"
	"github.com/Korpenter/assistand/internal/service"
	"github.com/Korpenter/assistand/internal/storage"
	"github.com/redis/go-redis/v9"
)

func main() {
	conf := flag.String("conf", "./config.json", "Path to the configuration file")
	flag.Parse()
	cfg := config.LoadConfig(*conf)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	hub := hub.NewMemHub()

	sessionStore := storage.NewRedisSessionStorage(redisClient, cfg.SessionDuration)
	sessionService := service.NewSessionService(sessionStore)
	matchService := service.NewMatchService()
	executeService := service.NewExecuteService()

	app := app.NewApp(cfg, sessionService, matchService, executeService, hub)
	app.Start()
}
