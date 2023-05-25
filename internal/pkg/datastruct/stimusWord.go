package datastruct

const StimusWordTableName = "stimus_word"

type StimusWord struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
