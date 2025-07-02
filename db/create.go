package db

import (
	"context"
	"fmt"
	"log"
	"strconv"
)

func Create[T Struct](table Table, obj *T) {
	v := getReflectValue(obj)

	var values []ColumnValue = parseValuesFromTable(table, v)

	colNames, valueNames, parsedValues := buildInsertParts(values)

	query := buildCreateQuery(table, colNames, valueNames)

	_, err := DB.Exec(context.Background(), query, parsedValues...)
	if err != nil {
		log.Fatal("failed to create object: ", err)
	}
}

func buildInsertParts(values []ColumnValue) (colNames string, valueNames string, parsedValues []any) {
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

func buildCreateQuery(table Table, colNames string, valueNames string) string {
	query := fmt.Sprintf(`
INSERT INTO %s (
%s
) VALUES (
%s
)
`, table.Name, colNames, valueNames)
	return query

}
