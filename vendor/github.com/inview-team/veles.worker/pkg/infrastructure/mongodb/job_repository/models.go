package job_repository

import (
	"github.com/inview-team/veles.worker/pkg/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Job struct {
	ID      primitive.ObjectID           `bson:"_id"`
	Name    string                       `bson:"name"`
	Actions []entities.ActionInformation `bson:"actions"`
	Output  entities.JobOutput           `bson:"output"`
}

func (j *Job) ToEntity() (*entities.Job, error) {
	return entities.NewJob(entities.JobID(j.ID.Hex()), j.Name, j.Actions, j.Output)
}

func NewJob(job *entities.Job) (*Job, error) {
	id, err := primitive.ObjectIDFromHex(string(job.Id))
	if err != nil {
		return nil, err
	}

	return &Job{
		ID:      id,
		Name:    job.Name,
		Actions: job.Actions,
		Output:  job.Output,
	}, nil
}
