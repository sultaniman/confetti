package repo

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"time"
)

type Repo struct {
	DB *sqlx.DB
	Q  *sq.StatementBuilderType
}

func (r *Repo) Select(table string) sq.SelectBuilder {
	return r.Q.Select("*").From(table)
}

func (r *Repo) Insert(table string, columns ...string) sq.InsertBuilder {
	return r.Q.Insert(table).Columns(columns...).Suffix("returning *")
}

func (r *Repo) Update(table string, includeUpdateDate bool) sq.UpdateBuilder {
	query := r.Q.Update(table)
	if includeUpdateDate {
		query = query.Set("updated_at", time.Now().UTC())
	}
	return query.Suffix("returning *")
}

func (r *Repo) Delete(table string, wheres sq.Eq) sq.DeleteBuilder {
	return r.Q.Delete(table).Where(wheres).Suffix("returning *")
}

func (r *Repo) Count(table string, wheres sq.Eq) sq.SelectBuilder {
	return r.Q.Select("COUNT(id)").From(table).Where(wheres).Suffix("returning *").Limit(1)
}
