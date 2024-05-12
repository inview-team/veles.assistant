package mongodb

import (
	"context"
	"fmt"

	"github.com/inview-team/veles.worker/pkg/infrastructure/mongodb/common"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
	*mongo.Client
}

type Config struct {
	IP         string `json:"ip"`
	Port       int    `json:"port"`
	User       string `json:"user"`
	Password   string `json:"password"`
	AuthSource string `json:"auth_source"`
}

func NewClient(ctx context.Context, cfg Config) (*DBClient, error) {
	URI := fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=%s", cfg.User, cfg.Password, cfg.IP, cfg.Port, cfg.AuthSource)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &DBClient{client}, nil
}

func (c *DBClient) Collection(dbName string, collName common.CollectionName) *mongo.Collection {
	return c.Database(dbName).Collection(string(collName))
}
