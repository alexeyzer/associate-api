package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
)

type StimusWordQuery interface {
	Create(ctx context.Context, req datastruct.StimusWord) (*datastruct.StimusWord, error)
	GetByID(ctx context.Context, ID int64) (*datastruct.StimusWord, error)
	GetByName(ctx context.Context, name string) (*datastruct.StimusWord, error)
	Exists(ctx context.Context, name string) (bool, error)
	List(ctx context.Context) ([]*datastruct.StimusWord, error)
}

type stimusWordQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *stimusWordQuery) List(ctx context.Context) ([]*datastruct.StimusWord, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.StimusWordTableName)
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var stimusWords []*datastruct.StimusWord

	err = q.db.SelectContext(ctx, &stimusWords, query, args...)
	if err != nil {
		return nil, err
	}

	return stimusWords, nil
}

func (q *stimusWordQuery) Create(ctx context.Context, req datastruct.StimusWord) (*datastruct.StimusWord, error) {
	qb := q.builder.Insert(datastruct.StimusWordTableName).
		Columns(
			"name",
		).
		Values(
			req.Name,
		).
		Suffix("ON CONFLICT DO NOTHING ").
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	fmt.Print(query)

	var stimusWord datastruct.StimusWord

	err = q.db.GetContext(ctx, &stimusWord, query, args...)
	if err != nil {
		return nil, err
	}

	return &stimusWord, nil
}

func (q *stimusWordQuery) GetByName(ctx context.Context, name string) (*datastruct.StimusWord, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.StimusWordTableName).
		Where(squirrel.Eq{"name": name})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var stimusWord datastruct.StimusWord

	err = q.db.GetContext(ctx, &stimusWord, query, args...)
	if err != nil {
		return nil, err
	}

	return &stimusWord, nil
}

func (q *stimusWordQuery) GetByID(ctx context.Context, ID int64) (*datastruct.StimusWord, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.StimusWordTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var stimusWord datastruct.StimusWord

	err = q.db.GetContext(ctx, &stimusWord, query, args...)
	if err != nil {
		return nil, err
	}

	return &stimusWord, nil
}

func (q *stimusWordQuery) Exists(ctx context.Context, name string) (bool, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.StimusWordTableName).
		Where(squirrel.Eq{"name": name})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var stimusWord datastruct.StimusWord

	err = q.db.GetContext(ctx, &stimusWord, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewStimusWordQuery(db *sqlx.DB) StimusWordQuery {
	return &stimusWordQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
