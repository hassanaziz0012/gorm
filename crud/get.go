package crud

import (
	"gorm/builder"
	"gorm/types"
)

func Get[T types.Struct](table types.Table[T], obj T, filters map[string]any) (items []T, err error) {
	return builder.
		Query(obj).
		Select().
		Where(filters).
		Limit(1).
		Build().
		Execute()
}
