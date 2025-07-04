package builder

import (
	types "gorm/types"

	"github.com/jackc/pgx/v5"
)

type QueryBuilder[T types.Struct] struct {
	model          T
	table          types.Table[T]
	parameterIndex int
	queryType      QueryType
	selectCols     []string
	update         map[string]any
	where          map[string]any
	orderBy        string
	orderDirection OrderDirection
	limit          int
	offset         int
	shouldReturn   bool
	returnedRows   pgx.Rows

	finalQuery  string
	finalValues []any
}

type QueryType int

const (
	SELECT QueryType = iota
	INSERT
	UPDATE
	DELETE
)

type OrderDirection string

const (
	Asc  OrderDirection = "asc"
	Desc OrderDirection = "desc"
)
