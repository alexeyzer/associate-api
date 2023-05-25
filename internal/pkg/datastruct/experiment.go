package datastruct

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

const ExperimentTableName = "experiment"

type ExperimentStatus string

const (
	ExperimentStatus_DONE        ExperimentStatus = "DONE"
	ExperimentStatus_CANCELED    ExperimentStatus = "CANCELED"
	ExperimentStatus_IN_PROGRESS ExperimentStatus = "IN_PROGRESS"
	ExperimentStatus_CREATED     ExperimentStatus = "CREATED"
)

type Experiment struct {
	ID                 int64              `db:"id"`
	Name               string             `db:"name"`
	Description        *string            `db:"description"`
	CreatorID          int64              `db:"creator_id"`
	RequiredAmount     *int64             `db:"required_amount"`
	ExperimentStimuses ExperimentStimuses `db:"experiment_stimuses"`
	Status             ExperimentStatus   `db:"status"`
}

type ExperimentStimus struct {
	StimusID            int64  `db:"stimus_id"`
	Name string
	LimitedResponseTime *int64 `db:"limited_response_time"`
}

type ExperimentStimuses []*ExperimentStimus

func (o ExperimentStimuses) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *ExperimentStimuses) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &o)
}
