package job_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/inview-team/veles.worker/pkg/domain/entities"
	"github.com/inview-team/veles.worker/pkg/infrastructure/mongodb"
	"github.com/inview-team/veles.worker/pkg/infrastructure/mongodb/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type jobRepository struct {
	client *mongodb.DBClient
	coll   *mongo.Collection
}

func NewJobRepository(client *mongodb.DBClient) entities.JobRepository {
	coll := client.Collection("assistant", common.JobCollectionName)
	return &jobRepository{
		client: client,
		coll:   coll,
	}
}

func (j jobRepository) Create(ctx context.Context, job *entities.Job) error {
	mJob, err := NewJob(job)
	if err != nil {
		return fmt.Errorf("creating action got error: %v", err)
	}

	if _, err := j.coll.InsertOne(ctx, mJob); err != nil {
		return fmt.Errorf("creating action got error: %v", err)
	}
	return nil
}

func (j jobRepository) GetByID(ctx context.Context, id string) (*entities.Job, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("finding action by id got err: %v", err)
	}

	f := bson.D{{"_id", oID}}
	res := j.coll.FindOne(ctx, f)

	var mJob Job
	if err = res.Decode(&mJob); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, entities.ErrActionNotFound
		}
		return nil, fmt.Errorf("finding action by id got error: %v", err)
	}

	return mJob.ToEntity()
}

func (j jobRepository) NextID(_ context.Context) entities.JobID {
	return entities.JobID(primitive.NewObjectID().Hex())
}
