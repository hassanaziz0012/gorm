package crud

import (
	"gorm/builder"
	"gorm/builder/filters"
	"gorm/types"
)

func Filter[T types.Struct](table types.Table[T], conditions []filters.ConditionGroup) (items []T, err error) {
	return builder.
		Query(table.Model).
		Select().
		Where(conditions).
		Build().
		Execute()
}
