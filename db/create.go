package db

import (
	"context"
	"fmt"
	"gorm/types"
	"log"
	"strconv"
)

func Create[T types.Struct](table types.Table[T], obj *T) {
	v := getReflectValue(obj)

	if err := ValidateObject(table, obj); err != nil {
		log.Fatal("failed to validate object: ", err)
	}

	var values []types.ColumnValue = parseValuesFromTable(table, v)

	colNames, valueNames, parsedValues := buildInsertParts(values)

	query := buildCreateQuery(table, colNames, valueNames)

	if _, err := DB.Exec(context.Background(), query, parsedValues...); err != nil {
		log.Fatal("failed to create object: ", err)
	}
}

func buildInsertParts(values []types.ColumnValue) (colNames string, valueNames string, parsedValues []any) {
	for i, value := range values {
		colNames += value.Colname
		valueNames += "$" + strconv.Itoa(i+1)
		if i+1 != len(values) {
			colNames += ", "
			valueNames += ", "
		}
		parsedValues = append(parsedValues, value.Value)
	}

	return colNames, valueNames, parsedValues
}

func buildCreateQuery[T any](table types.Table[T], colNames string, valueNames string) string {
	query := fmt.Sprintf(`
INSERT INTO %s (
%s
) VALUES (
%s
)
`, table.Name, colNames, valueNames)
	return query

}
