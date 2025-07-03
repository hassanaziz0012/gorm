package db

import (
	"context"
	"fmt"
	"gorm/types"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"
)

func Filter[T Struct](table types.Table, out *[]T, filters []types.ColumnValue) {
	filtersQuery, parsedValues := parseFilters(filters)
	query := buildFilterQuery(table, filtersQuery)

	rows, err := DB.Query(context.Background(), query, parsedValues...)
	if err != nil {
		log.Fatal("unable to filter rows: ", err)
	}
	defer rows.Close()

	parseRows(table, rows, out)

	if err := rows.Err(); err != nil {
		log.Fatal("row iteration error: ", err)
	}
}

func buildFilterQuery(table types.Table, filtersQuery string) string {
	query := fmt.Sprintf("SELECT * FROM %s WHERE", table.Name)

	query += filtersQuery

	return query
}

func parseRows[T Struct](table types.Table, rows pgx.Rows, out *[]T) {
	for rows.Next() {
		var item T
		v := reflect.ValueOf(&item).Elem()

		dest := prepareScanDest(table, v)

		if err := rows.Scan(dest...); err != nil {
			log.Fatal("failed to scan row: ", err)
		}

		*out = append(*out, item)
	}
}
