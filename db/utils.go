package db

import (
	"fmt"
	types "gorm/types"
	"reflect"
	"regexp"
	"strings"
)

// Given a table and an object `[v reflect.Value]`, extracts and returns a list of column names and their corresponding values.
func ParseValuesFromTable[T any](table types.Table[T], v reflect.Value) (values []types.ColumnValue) {
	for _, col := range table.Cols {
		if col.Constraints.AutoIncrement {
			continue
		}

		field := v.FieldByName(col.FieldName)
		if !field.IsValid() {
			panic(fmt.Sprintf("Column %s not found in struct %s", col.FieldName, v.Type().Name()))
		}

		value := field.Interface()
		values = append(values, types.ColumnValue{
			Colname: col.Name,
			Value:   value,
		})
	}
	return values
}

func PrepareScanDest[T any](table types.Table[T], v reflect.Value) (dest []any) {
	for _, col := range table.Cols {
		field := v.FieldByName(col.FieldName)
		if !field.IsValid() || !field.CanSet() {
			fmt.Println("field not found or unsettable: ", col.FieldName)
			continue
		}
		dest = append(dest, field.Addr().Interface())
	}
	return dest
}

func ToSnakeCase(s string) string {
	s = strings.TrimSpace(s)

	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = re.ReplaceAllString(s, "_")

	re2 := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	s = re2.ReplaceAllString(s, "${1}_${2}")

	re3 := regexp.MustCompile(`([A-Z]+)([A-Z][a-z])`)
	s = re3.ReplaceAllString(s, "${1}_${2}")

	s = strings.ToLower(s)

	s = strings.Trim(s, "_")
	s = regexp.MustCompile(`_+`).ReplaceAllString(s, "_")

	return s
}

func ParseDataType(typeName string) (types.DataType, error) {
	switch typeName {
	case "string":
		return types.String, nil
	case "uint", "int":
		return types.Integer, nil
	case "bool":
		return types.Boolean, nil
	case "Time":
		return types.Time, nil
	}

	return types.String, fmt.Errorf("invalid type name")
}
