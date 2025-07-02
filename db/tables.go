package db

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Column struct {
	// The struct name in Go
	FieldName string
	// The SQL table column name
	Name        string
	Coltype     DataType
	Constraints Constraints
}

type Constraints struct {
	DefaultNow    bool
	AutoIncrement bool
	IsPrimary     bool
	IsUnique      bool
}

type Table struct {
	Name string
	Cols []Column
}

type ColumnValue struct {
	Colname string
	Value   any
}

func ParseConstraints(tag reflect.StructTag) Constraints {
	c := Constraints{}
	for t := range strings.SplitSeq(tag.Get("gorm.constraints"), ",") {
		if t == "pk" {
			c.IsPrimary = true
		}
		if t == "autoincrement" {
			c.AutoIncrement = true
		}
		if t == "unique" {
			c.IsUnique = true
		}
	}

	defaultVal := tag.Get("gorm.default")
	if defaultVal == "now" {
		c.DefaultNow = true
	}

	return c
}

func parseCols(t reflect.Type) []Column {
	var cols []Column

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		fieldName := field.Name
		colname := toSnakeCase(field.Name)
		coltype, err := ParseDataType(field.Type.Name())
		if err != nil {
			log.Fatal("invalid data type: ", field.Type.Name())
		}
		constraints := ParseConstraints(field.Tag)

		col := Column{
			FieldName:   fieldName,
			Name:        colname,
			Coltype:     coltype,
			Constraints: constraints,
		}
		cols = append(cols, col)
	}
	return cols
}

// Converts any given struct into a Table object.
func createTableFromModel(model any) Table {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	cols := parseCols(t)

	table := Table{
		Name: strings.ToLower(t.Name()) + "s", // add an "s" to make it plural (e.g Struct "User", table name "users")
		Cols: cols,
	}
	return table
}

func buildColsQuery(table Table) (colQueries []string) {
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

		if col.Constraints.DefaultNow {
			query += " DEFAULT NOW()"
		}

		colQueries = append(colQueries, query)
	}
	return colQueries
}

func buildTableQuery(table Table, colQueries []string) string {
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

func CreateTable(model any) Table {
	table := createTableFromModel(model)
	colQueries := buildColsQuery(table)
	query := buildTableQuery(table, colQueries)
	fmt.Println(query)

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatal("failed to create table: ", err)
	}

	return table
}
