package crud

import (
	"gorm/builder"
	"gorm/types"
	"gorm/utils"
)

func Create[T types.Struct](table types.Table[T], obj *T) (items []T, err error) {
	values := utils.StructToColVals(table, obj, true)
	return builder.
		Query(table.Model).
		Insert(values).
		Build().
		Execute()
}
