package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
)

type ExperimentResultQuery interface {
	BatchCreate(ctx context.Context, req []*datastruct.ExperimentResult) ([]*datastruct.ExperimentResultResp, error)
	List(ctx context.Context, userID int64, experimentIDs []int64) ([]*datastruct.ExperimentResultList, error)
}

type experimentResultQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *experimentResultQuery) List(ctx context.Context, userID int64, experimentIDs []int64) ([]*datastruct.ExperimentResultList, error) {
	qb := q.builder.
		Select("ert.*, swt.name as stimus_word, awt.name as assotiation_word").
		From(datastruct.ExperimentResultTableName + " ert").
		Join(datastruct.StimusWordTableName + " swt on swt.id = ert.stimus_word_id").
		Join(datastruct.AssociateWordTableName + " awt on awt.id = ert.assotiation_word_id")

	if userID != 0 {
		qb = qb.Where(squirrel.Eq{"user_id": userID})
	}

	if len(experimentIDs) != 0 {
		qb = qb.Where(squirrel.Eq{"experiment_id": experimentIDs})
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var results []*datastruct.ExperimentResultList

	err = q.db.SelectContext(ctx, &results, query, args...)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (q *experimentResultQuery) BatchCreate(ctx context.Context, req []*datastruct.ExperimentResult) ([]*datastruct.ExperimentResultResp, error) {
	qb := q.builder.Insert(datastruct.ExperimentResultTableName).
		Columns(
			"id",
			"experiment_id",
			"user_id",
			"session_id",
			"stimus_word_id",
			"assotiation_word_id",
			"time_spend",
		)

	for _, r := range req {
		qb = qb.Values(
			r.ID,
			r.ExperimentID,
			r.UserID,
			r.SessionID,
			r.StimusWordID,
			r.AssotiationWordID,
			r.TimeSpend,
		)
	}
	qb = qb.Suffix("RETURNING *")

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var experimentsResult []*datastruct.ExperimentResultResp

	err = q.db.SelectContext(ctx, &experimentsResult, query, args...)
	if err != nil {
		return nil, err
	}

	return experimentsResult, nil
}

func NewExperimentResultQuery(db *sqlx.DB) ExperimentResultQuery {
	return &experimentResultQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
