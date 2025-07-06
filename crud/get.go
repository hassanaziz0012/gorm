package crud

import (
	"fmt"
	"gorm/builder"
	"gorm/builder/filters"
	"gorm/types"
)

func Get[T types.Struct](table types.Table[T], conditions []filters.ConditionGroup) (item T, err error) {
	items, err := builder.
		Query(table.Model).
		Select().
		Where(conditions).
		Limit(1).
		Build().
		Execute()

	if err != nil {
		return item, err
	}

	if len(items) == 0 {
		return item, fmt.Errorf("no results found")
	}

	return items[0], err
}
