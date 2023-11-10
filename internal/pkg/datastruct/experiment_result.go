package datastruct

const ExperimentResultTableName = "experiment_result"

type ExperimentResult struct {
	ID                int64 `db:"id"`
	ExperimentID      int64 `db:"experiment_id"`
	UserID            int64 `db:"user_id"`
	SessionID         int64 `db:"session_id"`
	StimusWordID      int64 `db:"stimus_word_id"`
	AssotiationWordID int64 `db:"assotiation_word_id"`
	AssotiationWord string
	TimeSpend         int64 `db:"time_spend"`
}

type ExperimentResultResp struct {
	ID                int64 `db:"id"`
	ExperimentID      int64 `db:"experiment_id"`
	UserID            int64 `db:"user_id"`
	SessionID         int64 `db:"session_id"`
	StimusWordID      int64 `db:"stimus_word_id"`
	AssotiationWordID int64 `db:"assotiation_word_id"`
	TimeSpend         int64 `db:"time_spend"`
}

type ExperimentResultList struct {
	ID                int64 `db:"id"`
	ExperimentID      int64 `db:"experiment_id"`
	UserID            int64 `db:"user_id"`
	SessionID         int64 `db:"session_id"`
	StimusWordID      int64 `db:"stimus_word_id"`
	StimusWord        string `db:"stimus_word"`
	AssotiationWordID int64 `db:"assotiation_word_id"`
	AssotiationWord string `db:"assotiation_word"`
	TimeSpend         int64 `db:"time_spend"`
}
