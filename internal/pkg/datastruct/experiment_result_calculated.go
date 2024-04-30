package datastruct

const ExperimentResultCalculatedTableName = "experiment_result_calculated"


type ExperimentResultCalculated struct {
	ExperimentID      int64 `db:"experiment_id"`
	StimusWordID      int64 `db:"stimus_word_id"`
	AssotiationWordID int64 `db:"assotiation_word_id"`
	Amount            int64 `db:"amount"`
	StimuesWord            string `db:"stimus_word"`
	AssotiationWord            string `db:"assotiation_word"`
}
