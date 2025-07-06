package builder

import (
	"gorm/builder/filters"
	"gorm/builder/tables"
)

func Query[T any](model T) *QueryBuilder[T] {
	table := tables.CreateTableFromModel(model)
	return &QueryBuilder[T]{model: model, table: table, parameterIndex: 1}
}

func (q *QueryBuilder[T]) Select(cols ...string) *QueryBuilder[T] {
	q.queryType = SELECT
	q.shouldReturn = true
	q.selectCols = cols
	return q
}

func (q *QueryBuilder[T]) Insert(values map[string]any) *QueryBuilder[T] {
	q.queryType = INSERT
	q.insert = values
	return q
}

func (q *QueryBuilder[T]) Update(values map[string]any) *QueryBuilder[T] {
	q.queryType = UPDATE
	q.update = values
	return q
}

func (q *QueryBuilder[T]) Delete() *QueryBuilder[T] {
	q.queryType = DELETE
	return q
}

func (q *QueryBuilder[T]) Where(conditions []filters.ConditionGroup) *QueryBuilder[T] {
	q.where = conditions
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
