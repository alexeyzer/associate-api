package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/alexeyzer/associate-api/config"
)

type DAO interface {
	UserQuery() UserQuery
	RoleQuery() RoleQuery
	UserRoleQuery() UserRoleQuery
	ExperimentQuery() ExperimentQuery
	StimusWordQuery() StimusWordQuery
}

type dao struct {
	userRoleQuery   UserRoleQuery
	roleQuery       RoleQuery
	userQuery       UserQuery
	experimentQuery ExperimentQuery
	stimusWordQuery StimusWordQuery
	db              *sqlx.DB
}

func NewDao() (DAO, error) {
	dao := &dao{}
	dbConf := config.Config.Database
	dsn := fmt.Sprintf(dbConf.Dsn,
		dbConf.Host,
		dbConf.Port,
		dbConf.Dbname,
		dbConf.User,
		dbConf.Password,
		dbConf.Ssl)
	fmt.Println(dsn)
	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	dao.db = conn
	return dao, nil
}

func (d *dao) RoleQuery() RoleQuery {
	if d.roleQuery == nil {
		d.roleQuery = NewRoleQuery(d.db)
	}
	return d.roleQuery
}

func (d *dao) UserRoleQuery() UserRoleQuery {
	if d.userRoleQuery == nil {
		d.userRoleQuery = NewUserRolesQuery(d.db)
	}
	return d.userRoleQuery
}

func (d *dao) UserQuery() UserQuery {
	if d.userQuery == nil {
		d.userQuery = NewUserQuery(d.db)
	}
	return d.userQuery
}

func (d *dao) ExperimentQuery() ExperimentQuery {
	if d.experimentQuery == nil {
		d.experimentQuery = NewExperimentQuery(d.db)
	}
	return d.experimentQuery
}

func (d *dao) StimusWordQuery() StimusWordQuery {
	if d.stimusWordQuery == nil {
		d.stimusWordQuery = NewStimusWordQuery(d.db)
	}
	return d.stimusWordQuery
}
