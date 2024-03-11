package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
)

type ExperimentQuery interface {
	Create(ctx context.Context, req datastruct.Experiment) (*datastruct.Experiment, error)
	GetByID(ctx context.Context, ID int64) (*datastruct.Experiment, error)
	Exists(ctx context.Context, name string) (bool, error)
	List(ctx context.Context, number, limit, userID int64, userExperiments bool, name *string) ([]*datastruct.Experiment, error)
}

type experimentQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *experimentQuery) List(ctx context.Context, number, limit, userID int64, userExperiments bool, name *string) ([]*datastruct.Experiment, error) {
	qb := q.builder.
		Select("*").
		From(datastruct.ExperimentTableName)

	if limit != 0 {
		qb = qb.Limit(uint64(limit))
	}
	if number != 0 {
		qb = qb.Offset(uint64((number - 1) * limit))
	}
	if userExperiments {
		qb = qb.Where(squirrel.Eq{"creator_id": userID})
	}

	if name != nil {
		qb = qb.Where(squirrel.Like{"LOWER(name)": fmt.Sprintf("%%%s%%", strings.ToLower(*name))})
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	print(query, number, limit)

	var experiments []*datastruct.Experiment

	err = q.db.SelectContext(ctx, &experiments, query, args...)
	if err != nil {
		return nil, err
	}

	return experiments, nil
}

func (q *experimentQuery) Create(ctx context.Context, req datastruct.Experiment) (*datastruct.Experiment, error) {
	qb := q.builder.Insert(datastruct.ExperimentTableName).
		Columns(
			"name",
			"creator_id",
			"description",
			"status",
			"required_amount",
			"experiment_stimuses",
			"conducducted_amount",
		).
		Values(
			req.Name,
			req.CreatorID,
			req.Description,
			req.Status,
			req.RequiredAmount,
			req.ExperimentStimuses,
			0,
		).
		Suffix("RETURNING *")
	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var experiment datastruct.Experiment

	err = q.db.GetContext(ctx, &experiment, query, args...)
	if err != nil {
		fmt.Print("jere")
		return nil, err
	}

	return &experiment, nil
}

func (q *experimentQuery) GetByID(ctx context.Context, ID int64) (*datastruct.Experiment, error) {
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

func (q *experimentQuery) Exists(ctx context.Context, name string) (bool, error) {
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

func NewExperimentQuery(db *sqlx.DB) ExperimentQuery {
	return &experimentQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
