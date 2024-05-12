package scenario_repository

import (
	"github.com/inview-team/veles.worker/pkg/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Scenario struct {
	ID           primitive.ObjectID `bson:"_id"`
	RootJobID    string             `bson:"root_job_id"`
	JobSequence  map[string]string  `bson:"job_sequence"`
	FinalMessage string             `bson:"final_message"`
}

func (s *Scenario) ToEntity() (*entities.Scenario, error) {
	jobSequence := make(map[entities.JobID]entities.JobID)
	for k, v := range s.JobSequence {
		jobSequence[entities.JobID(k)] = entities.JobID(v)
	}

	return entities.NewScenario(
		entities.ScenarioID(s.ID.Hex()),
		entities.JobID(s.RootJobID),
		jobSequence,
	)
}

func NewScenario(scenario *entities.Scenario) (*Scenario, error) {
	id, err := primitive.ObjectIDFromHex(string(scenario.ID))
	if err != nil {
		return nil, err
	}

	jobSequence := make(map[string]string)
	for k, v := range scenario.JobSequence {
		jobSequence[string(k)] = string(v)
	}

	return &Scenario{
		ID:           id,
		RootJobID:    string(scenario.RootJobID),
		JobSequence:  jobSequence,
		FinalMessage: scenario.FinalMessage,
	}, nil
}
