package tables

import (
	"context"
	"fmt"
	"gorm/db"
	"gorm/types"
	"log"
	"reflect"
	"strings"
)

func GenerateTableName(t reflect.Type) string {
	return strings.ToLower(t.Name()) + "s" // add an "s" to make it plural (e.g Struct "User", table name "users")
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

func buildColsQuery[T any](table types.Table[T]) (colQueries []string) {
	for _, col := range table.Cols {
		// as SERIAL is a data type, not a constraint, i used this hacky workaround.
		colType := col.Coltype.String()
		if col.Constraints.AutoIncrement {
			colType = "SERIAL"
		}

		query := fmt.Sprintf("%s %s", col.Name, colType)
		if col.Constraints.IsPrimary {
			query += " PRIMARY KEY"
		}

		if col.Constraints.IsUnique {
			query += " UNIQUE"
		}

		if col.Defaults.TimeNow {
			query += " DEFAULT NOW()"
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

func CreateTable[T any](model T) types.Table[T] {
	table := CreateTableFromModel(model)
	colQueries := buildColsQuery(table)
	query := buildTableQuery(table, colQueries)

	_, err := db.DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatal("failed to create table: ", err)
	}

	return table
}
