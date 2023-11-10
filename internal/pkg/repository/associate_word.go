package repository

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
)

type AssociateWordQuery interface {
	Create(ctx context.Context, req datastruct.AssociateWord) (*datastruct.AssociateWord, error)
	GetByID(ctx context.Context, ID int64) (*datastruct.AssociateWord, error)
	GetByName(ctx context.Context, name string) (*datastruct.AssociateWord, error)
	Exists(ctx context.Context, name string) (bool, error)
	List(ctx context.Context) ([]*datastruct.AssociateWord, error)
}

type associateWordQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *associateWordQuery) List(ctx context.Context) ([]*datastruct.AssociateWord, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.AssociateWordTableName)
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var stimusWords []*datastruct.AssociateWord

	err = q.db.SelectContext(ctx, &stimusWords, query, args...)
	if err != nil {
		return nil, err
	}

	return stimusWords, nil
}

func (q *associateWordQuery) Create(ctx context.Context, req datastruct.AssociateWord) (*datastruct.AssociateWord, error) {
	qb := q.builder.Insert(datastruct.AssociateWordTableName).
		Columns(
			"name",
		).
		Values(
			req.Name,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var stimusWord datastruct.AssociateWord

	err = q.db.GetContext(ctx, &stimusWord, query, args...)
	if err != nil {
		return nil, err
	}

	return &stimusWord, nil
}

func (q *associateWordQuery) GetByID(ctx context.Context, ID int64) (*datastruct.AssociateWord, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.AssociateWordTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var stimusWord datastruct.AssociateWord

	err = q.db.GetContext(ctx, &stimusWord, query, args...)
	if err != nil {
		return nil, err
	}

	return &stimusWord, nil
}

func (q *associateWordQuery) GetByName(ctx context.Context, name string) (*datastruct.AssociateWord, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.AssociateWordTableName).
		Where(squirrel.Eq{"name": name})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var stimusWord datastruct.AssociateWord

	err = q.db.GetContext(ctx, &stimusWord, query, args...)
	if err != nil {
		return nil, err
	}

	return &stimusWord, nil
}

func (q *associateWordQuery) Exists(ctx context.Context, name string) (bool, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.AssociateWordTableName).
		Where(squirrel.Eq{"name": name})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var stimusWord datastruct.AssociateWord

	err = q.db.GetContext(ctx, &stimusWord, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewAssociateWordQuery(db *sqlx.DB) AssociateWordQuery {
	return &associateWordQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
