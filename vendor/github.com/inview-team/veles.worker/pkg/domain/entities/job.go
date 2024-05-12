package entities

import "context"

type Job struct {
	Id      JobID
	Name    string
	Actions []ActionInformation
	Output  JobOutput
}

type JobID string
type JobType int

type JobOutput struct {
	Ask       Output
	OnSuccess Output
	OnFailure Output
}

type Output struct {
	Message  string
	Type     OutputType
	Variable map[string]Variable
}

type OutputType int

const (
	Ask OutputType = iota + 1
	Success
	Failure
)

type ActionInformation struct {
	Id     ActionID
	Output []string
}

type JobRepository interface {
	Create(ctx context.Context, job *Job) error
	GetByID(ctx context.Context, id string) (*Job, error)
	NextID(ctx context.Context) JobID
}

func NewJob(id JobID, name string, actions []ActionInformation, output JobOutput) (*Job, error) {
	return &Job{
		Id:      id,
		Name:    name,
		Actions: actions,
		Output:  output,
	}, nil
}
