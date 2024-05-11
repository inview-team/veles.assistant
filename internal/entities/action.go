package entities

type ActionType int

const (
	HTTP ActionType = iota + 1
	Condition
)

type Action struct {
	ID             string     `bson:"_id,omitempty"`
	Type           ActionType `bson:"type"`
	Next           NextAction `bson:"next"`
	InputTemplate  string     `bson:"input_template"`
	OutputTemplate string     `bson:"output_template"`
}

type NextAction struct {
	OnSuccess string `bson:"on_success"`
	OnFailure string `bson:"on_failure"`
}
