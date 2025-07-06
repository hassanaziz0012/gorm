package tables

import (
	"context"
	"fmt"
	"gorm/db"
	"gorm/types"
	"gorm/utils"
	"log"
	"reflect"
	"strings"
)

type TableBuilder[T any] struct {
	model  T
	table  types.Table[T]
	name   string
	query  string
	checks []Check
}

func GenerateTableName(t reflect.Type) string {
	return strings.ToLower(t.Name()) + "s" // add an "s" to make it plural (e.g Struct "User", table name "users")
}

func Table[T any](model T) *TableBuilder[T] {
	return &TableBuilder[T]{
		model: model,
		name:  GenerateTableName(reflect.TypeOf(model)),
	}
}

func (t *TableBuilder[T]) AddCheck(check Check) *TableBuilder[T] {
	t.checks = append(t.checks, check)
	return t
}

func (t *TableBuilder[T]) BuildQuery() *TableBuilder[T] {
	table := CreateTableFromModel(t.model)
	t.table = table

	colQueries := t.buildColsQuery()
	query := utils.RemoveExtraSpaces(buildTableQuery(t.table, colQueries))
	t.query = query
	return t
}

func (t *TableBuilder[T]) Execute() types.Table[T] {
	fmt.Println(t.query)
	_, err := db.DB.Exec(context.Background(), t.query)
	if err != nil {
		log.Fatal("failed to create table: ", err)
	}
	return t.table
}

// Converts any given struct into a Table object.
func CreateTableFromModel[T any](model T) types.Table[T] {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	cols := parseCols(t)

	table := types.Table[T]{
		Model: model,
		Name:  GenerateTableName(t),
		Cols:  cols,
	}
	return table
}

func buildColConstraints(col *types.Column) string {
	var query string
	if col.Constraints.Required && !col.Constraints.AutoIncrement {
		query += " NOT NULL "
	}
	if col.Constraints.IsUnique && !col.Constraints.AutoIncrement {
		query += " UNIQUE "
	}
	if col.Constraints.IsPrimary {
		query += " PRIMARY KEY "
	}

	return query
}

func buildColDefaults(col *types.Column) string {
	var query string
	if col.Defaults.TimeNow {
		query += " DEFAULT NOW() "
	}

	if col.Defaults.Bool == types.TRUE {
		query += " DEFAULT TRUE "
	}

	if col.Defaults.Bool == types.FALSE {
		query += " DEFAULT FALSE "
	}

	if !utils.IsEmpty(col.Defaults.Int) {
		query += fmt.Sprintf(" DEFAULT %d ", col.Defaults.Int)
	}

	if !utils.IsEmpty(col.Defaults.Text) {
		query += fmt.Sprintf(" DEFAULT '%s' ", col.Defaults.Text)
	}

	return query
}

func buildColType(col *types.Column) (colType string) {
	// as SERIAL is a data type, not a constraint, i used this hacky workaround.
	colType, _ = col.Coltype.String()
	if col.Constraints.AutoIncrement {
		colType = "SERIAL"
	}
	return colType
}

func (t *TableBuilder[T]) buildColsQuery() (colQueries []string) {
	for _, col := range t.table.Cols {
		var query string
		colType := buildColType(&col)
		query = fmt.Sprintf("%s %s", col.Name, colType)

		query += buildColConstraints(&col)
		query += buildColDefaults(&col)

		for _, check := range t.checks {
			if col.Name == check.Col {
				query += check.BuildClause()
			}
		}
		if !utils.IsEmpty(col.FKR) {
			constraintName := col.FKR.ConstraintName
			fieldName := col.FKR.Name
			tableName := col.FKR.FKTable
			fkField := col.FKR.FKField

			fk_query := fmt.Sprintf("CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(%s) ON DELETE CASCADE", constraintName, fieldName, tableName, fkField)
			colQueries = append(colQueries, fk_query)
		}

		colQueries = append(colQueries, query)
	}
	return colQueries
}

func buildTableQuery[T any](table types.Table[T], colQueries []string) string {
	query := fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
`, table.Name)
	for i, colQuery := range colQueries {
		if i+1 == len(colQueries) {
			query += colQuery + "\n"
		} else {
			query += colQuery + ",\n"
		}
	}
	query += ");"
	return query
}
