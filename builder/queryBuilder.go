package builder

import (
	"context"
	"fmt"
	"gorm/db"
	"gorm/tables"

	"github.com/jackc/pgx/v5"
)

func Query[T any](model T) *QueryBuilder[T] {
	table := tables.CreateTableFromModel(model)
	return &QueryBuilder[T]{model: model, table: table, parameterIndex: 1}
}

func (q *QueryBuilder[T]) Select(cols ...string) *QueryBuilder[T] {
	q.queryType = SELECT
	q.shouldReturn = true
	q.selectCols = append(q.selectCols, cols...)
	return q
}

func (q *QueryBuilder[T]) Update(values map[string]any) *QueryBuilder[T] {
	q.queryType = UPDATE
	q.update = values
	return q
}

func (q *QueryBuilder[T]) Where(filters map[string]any) *QueryBuilder[T] {
	q.where = filters
	return q
}

func (q *QueryBuilder[T]) OrderBy(col string, direction OrderDirection) *QueryBuilder[T] {
	q.orderBy = col
	q.orderDirection = direction
	return q
}

func (q *QueryBuilder[T]) Limit(lim int) *QueryBuilder[T] {
	q.limit = lim
	return q
}

func (q *QueryBuilder[T]) Offset(offset int) *QueryBuilder[T] {
	q.offset = offset
	return q
}

func (q *QueryBuilder[T]) Build() *QueryBuilder[T] {
	var query string
	var values []any

	switch q.queryType {
	case SELECT:
		query = q.buildSelect()
	case UPDATE:
		query, values = q.buildUpdate()
	}

	if len(q.where) > 0 {
		whereClause, whereValues := q.buildWhereClause()
		query += whereClause
		values = append(values, whereValues...)
	}

	if q.orderBy != "" {
		orderByClause := q.buildOrderClause()
		query += orderByClause
	}

	if q.limit > 0 {
		query += fmt.Sprintf(" LIMIT %d ", q.limit)
	}

	if q.offset > 0 {
		query += fmt.Sprintf(" OFFSET %d ", q.offset)
	}

	query = removeExtraSpaces(query)
	query += ";"

	q.finalQuery = query
	q.finalValues = values
	return q
}

func (q *QueryBuilder[T]) Execute() (items []T, err error) {
	fmt.Println(q.finalQuery, q.finalValues)
	if q.finalQuery == "" {
		return nil, fmt.Errorf("please Build() the query first")
	}

	var rows pgx.Rows
	if q.shouldReturn {
		rows, err = db.DB.Query(context.Background(), q.finalQuery, q.finalValues...)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		items = q.scanRows(rows)
		return items, err

	} else {
		_, err = db.DB.Exec(context.Background(), q.finalQuery, q.finalValues...)
		if err != nil {
			return nil, err
		}
	}

	return items, err
}
