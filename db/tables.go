package db

import (
	"context"
	"fmt"
	"gorm/types"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func parseConstraints(tag reflect.StructTag) types.Constraints {
	c := types.Constraints{}
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

func parseValidators(tag reflect.StructTag) types.Validators {
	var v types.Validators
	for val := range strings.SplitSeq(tag.Get("gorm.validators"), ",") {
		if val == "email" {
			v.IsEmail = true
		}
		if val == "url" {
			v.IsURL = true
		}
		if strings.HasPrefix(val, "min(") && strings.HasSuffix(val, ")") {
			extracted := strings.TrimSuffix(strings.TrimPrefix(val, "min("), ")")
			minLength, err := strconv.Atoi(extracted)
			if err != nil {
				log.Fatal("unable to parse min() value: ", err)
			}
			v.MinLength = minLength
		}
		if strings.HasPrefix(val, "max(") && strings.HasSuffix(val, ")") {
			extracted := strings.TrimSuffix(strings.TrimPrefix(val, "max("), ")")
			maxLength, err := strconv.Atoi(extracted)
			if err != nil {
				log.Fatal("unable to parse max() value: ", err)
			}
			v.MaxLength = maxLength
		}
	}
	return v
}

func parseCols(t reflect.Type) []types.Column {
	var cols []types.Column

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		fieldName := field.Name
		colname := ToSnakeCase(field.Name)
		coltype, err := ParseDataType(field.Type.Name())
		if err != nil {
			log.Fatal("invalid data type: ", field.Type.Name())
		}
		constraints := parseConstraints(field.Tag)
		validators := parseValidators(field.Tag)

		col := types.Column{
			FieldName:   fieldName,
			Name:        colname,
			Coltype:     coltype,
			Constraints: constraints,
			Validators:  validators,
		}
		cols = append(cols, col)
	}
	return cols
}

// Converts any given struct into a Table object.
func createTableFromModel(model any) types.Table {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	cols := parseCols(t)

	table := types.Table{
		Name: strings.ToLower(t.Name()) + "s", // add an "s" to make it plural (e.g Struct "User", table name "users")
		Cols: cols,
	}
	return table
}

func buildColsQuery(table types.Table) (colQueries []string) {
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

func buildTableQuery(table types.Table, colQueries []string) string {
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

func CreateTable(model any) types.Table {
	table := createTableFromModel(model)
	colQueries := buildColsQuery(table)
	query := buildTableQuery(table, colQueries)

	_, err := DB.Exec(context.Background(), query)
	if err != nil {
		log.Fatal("failed to create table: ", err)
	}

	return table
}
