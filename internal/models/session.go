package models

type Session struct {
	ID         string `json:"id"`
	Token      string `json:"token"`
	ScenarioID string `json:"scenario_id"`
	JobID      string `json:"job_id"`
}
