package tables

import (
	"context"
	"fmt"
	"gorm/db"
	"gorm/types"
)

func Exists(tablename string) bool {
	query := fmt.Sprintf("SELECT to_regclass('public.%s')", tablename)
	row := db.DB.QueryRow(context.Background(), query)

	var result string
	row.Scan(&result)

	return result != ""
}

func (t *TableBuilder[T]) GetExisting() (types.Table[T], error) {
	if Exists(t.name) {
		return CreateTableFromModel(t.model), nil
	} else {
		return types.Table[T]{}, fmt.Errorf("table does not exist")
	}
}
