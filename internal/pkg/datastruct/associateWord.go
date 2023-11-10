package datastruct

const AssociateWordTableName = "associate_word"

type AssociateWord struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
