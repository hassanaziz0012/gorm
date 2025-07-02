package db

import (
	"context"
	"fmt"
	"log"
)

func Get[T Struct](table Table, obj *T, filters []ColumnValue) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE", table.Name)

	filtersQuery, parsedValues := parseFilters(filters)
	query += filtersQuery

	v := getReflectValue(obj)

	dest := prepareScanDest(table, v)

	row := DB.QueryRow(context.Background(), query, parsedValues...)
	err := row.Scan(dest...)
	if err != nil {
		log.Fatal("unable to find object: ", err)
	}
}
