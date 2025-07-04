package crud

import (
	"gorm/builder"
	"gorm/types"
)

func Filter[T types.Struct](table types.Table[T], filters map[string]any) (items []T, err error) {
	return builder.Query(table.Model).Select().Where(filters).Build().Execute()

	// filtersQuery, parsedValues := parseFilters(filters)
	// query := buildFilterQuery(table, filtersQuery)

	// rows, err := db.DB.Query(context.Background(), query, parsedValues...)
	// if err != nil {
	// 	log.Fatal("unable to filter rows: ", err)
	// }
	// defer rows.Close()

	// parseRows(table, rows, out)

	// if err := rows.Err(); err != nil {
	// 	log.Fatal("row iteration error: ", err)
	// }
}

// func buildFilterQuery(table types.Table, filtersQuery string) string {
// 	query := fmt.Sprintf("SELECT * FROM %s WHERE", table.Name)

// 	query += filtersQuery

// 	return query
// }

// func parseRows[T types.Struct](table types.Table, rows pgx.Rows, out *[]T) {
// 	for rows.Next() {
// 		var item T
// 		v := reflect.ValueOf(&item).Elem()

// 		dest := db.PrepareScanDest(table, v)

// 		if err := rows.Scan(dest...); err != nil {
// 			log.Fatal("failed to scan row: ", err)
// 		}

// 		*out = append(*out, item)
// 	}
// }
