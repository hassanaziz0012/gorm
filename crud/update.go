package crud

import (
	"fmt"
	"gorm/builder"
	"gorm/builder/filters"
	"gorm/types"
	"gorm/utils"
	"reflect"
)

func Update[T types.Struct](table types.Table[T], obj *T) (items []T, err error) {
	var values map[string]any = utils.StructToColVals(table, obj, true)

	id := reflect.ValueOf(obj).Elem().FieldByName("ID").Uint()
	if utils.IsEmpty(id) {
		return items, fmt.Errorf("id is empty")
	}

	queryFilters := []filters.Filter{
		filters.EqualTo("id", int(id)),
	}
	conditions := []filters.ConditionGroup{
		{
			Filters: queryFilters,
			Combine: filters.AND,
		},
	}
	return builder.
		Query(table.Model).
		Update(values).
		Where(conditions).
		Build().
		Execute()
}
