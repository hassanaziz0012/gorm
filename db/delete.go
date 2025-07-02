package db

import (
	"context"
	"fmt"
)

func Delete[T Struct](table Table, obj T) {
	pk := getPrimaryKeyCol(table)
	v := getReflectValue(obj)
	id := v.FieldByName(pk.FieldName).Interface()

	query := buildDeleteQuery(table, pk)
	res, err := DB.Exec(context.Background(), query, id)
	if err != nil {
		if res.RowsAffected() == 0 {
			fmt.Println("no rows affected")
		}
	}
}

func getPrimaryKeyCol(table Table) Column {
	var pk Column
	for _, col := range table.Cols {
		if col.Constraints.IsPrimary {
			pk = col
		}
	}
	return pk
}

func buildDeleteQuery(table Table, pk Column) string {
	query := fmt.Sprintf(`
DELETE FROM %s WHERE %s = $1
`, table.Name, pk.FieldName)

	return query

}
