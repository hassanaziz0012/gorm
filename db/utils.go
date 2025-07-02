package db

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func getReflectValue(obj any) reflect.Value {
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("obj must be a struct or a pointer to a struct")
	}

	return v
}

// Converts the given list of filters to SQL "AND" conditions that can then be appended to a WHERE clause.
func parseFilters(filters []ColumnValue) (string, []any) {
	var parsedValues []any
	var query string
	for i, value := range filters {
		query += fmt.Sprintf(" %s = %s", value.Colname, "$"+strconv.Itoa(i+1))
		if i+1 != len(filters) {
			query += " AND"
		}
		parsedValues = append(parsedValues, value.Value)
	}
	return query, parsedValues
}

// Given a table and an object `[v reflect.Value]`, extracts and returns a list of column names and their corresponding values.
func parseValuesFromTable(table Table, v reflect.Value) (values []ColumnValue) {
	for _, col := range table.Cols {
		if col.Constraints.AutoIncrement {
			continue
		}

		field := v.FieldByName(col.FieldName)
		if !field.IsValid() {
			panic(fmt.Sprintf("Column %s not found in struct %s", col.FieldName, v.Type().Name()))
		}

		value := field.Interface()
		values = append(values, ColumnValue{
			Colname: col.Name,
			Value:   value,
		})
	}
	return values
}

func prepareScanDest(table Table, v reflect.Value) (dest []any) {
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

func toSnakeCase(s string) string {
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
