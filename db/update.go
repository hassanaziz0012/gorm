package db

import (
	"context"
	"fmt"
	"gorm/types"
	"log"
	"reflect"
	"strconv"
)

func Update[T types.Struct](table types.Table[T], obj *T) {
	v := getReflectValue(obj)

	values := parseValuesFromTable(table, v)

	fieldQueries, parsedValues := buildSetClause(values)

	id := extractPrimaryKey(table, v)
	parsedValues = append(parsedValues, id)

	query := buildUpdateQuery(table, fieldQueries, values)

	dest := PrepareScanDest(table, v)

	err := DB.QueryRow(context.Background(), query, parsedValues...).Scan(dest...)
	if err != nil {
		log.Fatal("unable to update object: ", err)
	}
}

func buildSetClause(values []types.ColumnValue) (fieldQueries string, parsedValues []any) {
	for i, val := range values {
		fieldQuery := fmt.Sprintf(" %s = %s", val.Colname, "$"+strconv.Itoa(i+1))
		if i+1 != len(values) {
			fieldQuery += ", "
		}
		fieldQueries += fieldQuery
		parsedValues = append(parsedValues, val.Value)
	}
	return fieldQueries, parsedValues
}

func extractPrimaryKey[T any](table types.Table[T], v reflect.Value) any {
	var id any
	for _, col := range table.Cols {
		if col.Constraints.IsPrimary {
			id = v.FieldByName(col.FieldName).Interface()
		}
	}
	return id
}

func buildUpdateQuery[T any](table types.Table[T], fieldQueries string, values []types.ColumnValue) (query string) {
	idPlaceholder := "$" + strconv.Itoa(len(values)+1)
	query = fmt.Sprintf(`
UPDATE %s 
SET %s 
WHERE id = %s 
RETURNING *
	`, table.Name, fieldQueries, idPlaceholder)

	return query
}
