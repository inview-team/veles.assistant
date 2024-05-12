package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/inview-team/veles.assistant/internal/models"
	"github.com/inview-team/veles.assistant/pkg/proto/gen/pb"
	"github.com/inview-team/veles.worker/pkg/domain/entities"
	"github.com/inview-team/veles.worker/pkg/domain/usecases/job_usecases"
	"google.golang.org/grpc"
)

type ActionService interface {
	ProcessMessage(ctx context.Context, session *models.Session, text string) (string, string, string, error)
}

type actionService struct {
	scenarioRepository entities.ScenarioRepository
	actionStorage      entities.ActionRepository
	jobStorage         entities.JobRepository
	jobExecutor        job_usecases.JobUsecases
	grpcClient         pb.NLPServiceClient
}

func NewActionService(as entities.ActionRepository, js entities.JobRepository, sr entities.ScenarioRepository, e job_usecases.JobUsecases, grpcConn *grpc.ClientConn) ActionService {
	return &actionService{
		actionStorage:      as,
		jobStorage:         js,
		scenarioRepository: sr,
		jobExecutor:        e,
		grpcClient:         pb.NewNLPServiceClient(grpcConn),
	}
}

func (as *actionService) ProcessMessage(ctx context.Context, session *models.Session, text string) (string, string, string, error) {
	args := make(map[string]entities.Variable)
	var jobID, scenarioID string
	session.ScenarioID = ""
	if session.ScenarioID == "" {
		req := &pb.MatchScenarioRequest{
			UserPrompt: text,
		}
		res, err := as.grpcClient.MatchScenario(ctx, req)
		if err != nil {
			return "", "", "", fmt.Errorf("failed to match scenario: %v", err)
		}

		if res.RootId == "" {
			return "", "", "", fmt.Errorf("no matching scenario found")
		}

		scenario, err := as.scenarioRepository.GetById(ctx, res.RootId)
		if err != nil {
			return "", "", "", fmt.Errorf("failed to get scenario: %v", err)
		}

		session.ScenarioID = res.RootId
		session.JobID = string(scenario.RootJobID)
		jobID = session.JobID
		scenarioID = session.ScenarioID
	} else {
		jobID = session.JobID
		scenarioID = session.ScenarioID
	}

	extractArgsReq := &pb.ExtractArgumentsRequest{
		UserPrompt: text,
	}
	extractArgsRes, err := as.grpcClient.ExtractArguments(ctx, extractArgsReq)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to extract arguments: %v", err)
	}

	for key, value := range extractArgsRes.Arguments {
		args[key] = entities.Variable{Value: value}
	}

	args["token"] = entities.Variable{Value: session.Token}

	out, err := as.jobExecutor.Run(ctx, jobID, args)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to run job: %v", err)
	}

	tmpl, err := template.New("output").Parse(out.Message)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to parse output template: %v", err)
	}

	templateData := make(map[string]interface{})
	for key, variable := range out.Variable {
		templateData[key] = variable.Value
	}

	var result bytes.Buffer
	if err := tmpl.Execute(&result, templateData); err != nil {
		return "", "", "", fmt.Errorf("failed to execute template: %v", err)
	}

	scenario, err := as.scenarioRepository.GetById(ctx, scenarioID)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to get scenario: %v", err)
	}

	nextJobID, hasNextJob := scenario.JobSequence[entities.JobID(jobID)]
	if hasNextJob {
		session.JobID = string(nextJobID)
		session.ScenarioID = ""
		session.JobID = ""
	} else {
		session.ScenarioID = ""
		session.JobID = ""
	}

	return result.String(), "", "", nil
}
