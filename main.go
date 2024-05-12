package main

import (
	"context"
	"flag"

	"github.com/inview-team/veles.assistant/internal/app"
	"github.com/inview-team/veles.assistant/internal/config"
	"github.com/inview-team/veles.assistant/internal/hub"
	"github.com/inview-team/veles.assistant/internal/service"
	"github.com/inview-team/veles.assistant/internal/storage"
	"github.com/inview-team/veles.worker/pkg/domain/usecases/job_usecases"
	"github.com/inview-team/veles.worker/pkg/infrastructure/mongodb"
	"github.com/inview-team/veles.worker/pkg/infrastructure/mongodb/action_repository"
	"github.com/inview-team/veles.worker/pkg/infrastructure/mongodb/job_repository"
	"github.com/inview-team/veles.worker/pkg/infrastructure/mongodb/scenario_repository"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
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

	mongoClient, err := mongodb.NewClient(context.Background(), mongodb.Config{
		IP:         cfg.MongoIP,
		Port:       cfg.MongoPort,
		User:       cfg.MongoUser,
		Password:   cfg.MongoPassword,
		AuthSource: cfg.MongoAuthSource,
	})
	if err != nil {
		log.Fatalf("Failed to connect to mongoDB: %v", err)
	}

	actionStorage := action_repository.NewActionRepository(mongoClient)
	jobRepo := job_repository.NewJobRepository(mongoClient)
	scenarioRepository := scenario_repository.NewScenarioRepository(mongoClient)
	executor := job_usecases.New(jobRepo, actionStorage)
	actionService := service.NewActionService(actionStorage, jobRepo, scenarioRepository, executor, conn)

	app := app.NewApp(cfg, sessionService, actionService, hub)
	app.Start()
}
