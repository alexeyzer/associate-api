package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	"github.com/jmoiron/sqlx"
	"github.com/ovadbar/squirrel"
)

type ExperimentResultCalculatedQuery interface {
	BatchCreate(ctx context.Context, req []*datastruct.ExperimentResultCalculated) error
	GetCalculatedResulsts(ctx context.Context, userID int64, experimentIDs []int64, number, limit int64, names []string) ([]*datastruct.ExperimentResultCalculated, error)
}

type experimentResultCalculatedQuery struct {
	builder squirrel.StatementBuilderType
	db      *sqlx.DB
}

func (q *experimentResultCalculatedQuery) GetCalculatedResulsts(ctx context.Context, userID int64, experimentIDs []int64, number, limit int64, names []string) ([]*datastruct.ExperimentResultCalculated, error) {
	qb := q.builder.
		Select("ert.*, swt.name as stimus_word, awt.name as assotiation_word").
		From(datastruct.ExperimentResultCalculatedTableName + " ert").
		Join(datastruct.StimusWordTableName + " swt on swt.id = ert.stimus_word_id").
		Join(datastruct.AssociateWordTableName + " awt on awt.id = ert.assotiation_word_id")

	if userID != 0 {
		qb = qb.Where(squirrel.Eq{"user_id": userID})
	}

	if len(experimentIDs) != 0 {
		qb = qb.Where(squirrel.Eq{"experiment_id": experimentIDs})
	}



	if len(names) > 0 {
		massivee :=make([]string, 0)

		
		for _, name := range names {
			massivee = append(massivee,
				strings.ToLower(name))
		}

		qb = qb.Where(
			squirrel.Or{
				squirrel.Eq{"LOWER(swt.name)": massivee},
				squirrel.Eq{"LOWER(awt.name)": massivee},
			})
			
		qb = qb.Suffix("UNION select ert.*, swt.name as stimus_word, awt.name as assotiation_word from experiment_result_calculated ert join stimus_word swt on swt.id = ert.stimus_word_id join associate_word awt on awt.id = ert.assotiation_word_id where (swt.name in (awt.name) or awt.name in (swt.name)) and swt.name != awt.name")
		qb = q.builder.Select("*").From("result").WithRecursive("result", qb)
		query, args, err := qb.ToSql()
		if err != nil {
			return nil, err
		}
		fmt.Println(query, args)

		var results []*datastruct.ExperimentResultCalculated

		err = q.db.SelectContext(ctx, &results, query, args...)
		if err != nil {
			return nil, err
		}

		return results, nil
	}

	if limit != 0 {
		qb = qb.Limit(uint64(limit))
	}
	if number != 0 {
		qb = qb.Offset(uint64((number - 1) * limit))
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var results []*datastruct.ExperimentResultCalculated

	err = q.db.SelectContext(ctx, &results, query, args...)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (q *experimentResultCalculatedQuery) BatchCreate(ctx context.Context, req []*datastruct.ExperimentResultCalculated) error {
	qb := q.builder.Insert(datastruct.ExperimentResultCalculatedTableName).
		Columns(
			"experiment_id",
			"stimus_word_id",
			"assotiation_word_id",
			"amount",
		)

	i := 0
	for _, r := range req {
		qb = qb.Values(
			r.ExperimentID,
			r.StimusWordID,
			r.AssotiationWordID,
			r.Amount,
		)
		i++
		if i > 6000 {
			i = 0
			query, args, err := qb.ToSql()
			if err != nil {
				return err
			}
			_, err = q.db.ExecContext(ctx, query, args...)
			if err != nil {
				return err
			}
			qb = q.builder.Insert(datastruct.ExperimentResultCalculatedTableName).
				Columns(
					"experiment_id",
					"stimus_word_id",
					"assotiation_word_id",
					"amount",
				)

		}
	}
	qb = qb.Suffix("RETURNING *")

	query, args, err := qb.ToSql()
	if err != nil {
		return err
	}

	_, err = q.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func NewExperimentResultCalculatedQuery(db *sqlx.DB) ExperimentResultCalculatedQuery {
	return &experimentResultCalculatedQuery{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
