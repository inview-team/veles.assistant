package main

import (
	"context"
	"flag"
	"time"

	"github.com/inview-team/veles.assistant/internal/app"
	"github.com/inview-team/veles.assistant/internal/config"
	"github.com/inview-team/veles.assistant/internal/hub"
	"github.com/inview-team/veles.assistant/internal/service"
	"github.com/inview-team/veles.assistant/internal/storage"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	conn, err := grpc.Dial(cfg.MatchServiceGRPCAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to match service gRPC server: %v", err)
	}
	defer conn.Close()

	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := mongoClient.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	database := mongoClient.Database(cfg.MongoDBName)
	actionStorage := storage.NewMongoActionStorage(database, cfg.MongoActionCollection)

	actionService := service.NewActionService(actionStorage, conn)

	executeService := service.NewExecuteService()

	app := app.NewApp(cfg, sessionService, actionService, executeService, hub)
	app.Start()
}
