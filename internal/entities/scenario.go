package entities

type Scenario struct {
	ID       string   `bson:"_id,omitempty"`
	Keywords []string `bson:"keywords"`
	Example  string   `bson:"example"`
	RootID   string   `bson:"root_id"`
}
