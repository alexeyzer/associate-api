package repository

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
)

type ExperimentResultQuery interface {
	Create(ctx context.Context, req datastruct.Experiment) (*datastruct.Experiment, error)
	GetByID(ctx context.Context, ID int64) (*datastruct.Experiment, error)
	Exists(ctx context.Context, name string) (bool, error)
	List(ctx context.Context) ([]*datastruct.Experiment, error)
}

type experimentResultQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *experimentResultQuery) List(ctx context.Context) ([]*datastruct.Experiment, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.ExperimentTableName)
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var experiments []*datastruct.Experiment

	err = q.db.SelectContext(ctx, &experiments, query, args...)
	if err != nil {
		return nil, err
	}

	return experiments, nil
}

func (q *experimentResultQuery) Create(ctx context.Context, req datastruct.Experiment) (*datastruct.Experiment, error) {
	qb := q.builder.Insert(datastruct.ExperimentTableName).
		Columns(
			"name",
			"creator_id",
			"description",
			"status",
			"required_amount",
			"experiment_stimuses",
		).
		Values(
			req.Name,
			req.CreatorID,
			req.Description,
			req.Status,
			req.RequiredAmount,
			req.ExperimentStimuses,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var experiment datastruct.Experiment

	err = q.db.GetContext(ctx, &experiment, query, args...)
	if err != nil {
		return nil, err
	}

	return &experiment, nil
}

func (q *experimentResultQuery) GetByID(ctx context.Context, ID int64) (*datastruct.Experiment, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.ExperimentTableName).
		Where(squirrel.Eq{"id": ID})
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var experiment datastruct.Experiment

	err = q.db.GetContext(ctx, &experiment, query, args...)
	if err != nil {
		return nil, err
	}

	return &experiment, nil
}

func (q *experimentResultQuery) Exists(ctx context.Context, name string) (bool, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.ExperimentTableName).
		Where(squirrel.Eq{"name": name})
	query, args, err := qb.ToSql()
	if err != nil {
		return false, err
	}

	var experiment datastruct.Experiment

	err = q.db.GetContext(ctx, &experiment, query, args...)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func NewExperimentResultQuery(db *sqlx.DB) ExperimentResultQuery {
	return &experimentResultQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
