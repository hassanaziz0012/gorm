package utils

import (
	"gorm/types"
	"reflect"
	"regexp"
	"strings"
)

func GetReflectValue(obj any) reflect.Value {
	v := reflect.ValueOf(obj)

	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("obj must be a struct or a pointer to a struct")
	}

	return v
}

func GetReflectType(obj any) reflect.Type {
	t := reflect.TypeOf(obj)

	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	return t
}

func StructToColVals[T any](table types.Table[T], obj *T, skipPK bool) map[string]any {
	result := make(map[string]any)

	v := GetReflectValue(obj)

	for _, col := range table.Cols {
		value := v.FieldByName(col.FieldName)

		if skipPK && col.Name == "id" {
			continue
		}

		if !IsEmpty(col.FKR) {
			fkObjID := value.FieldByName("ID").Uint()
			result[col.Name] = fkObjID
			continue
		}

		if !value.CanInterface() || IsEmpty(value.Interface()) {
			continue
		}

		result[col.Name] = value.Interface()

	}

	return result
}

func IsEmpty(value any) bool {
	if value == nil {
		return true
	}

	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		return v.Len() == 0

	case reflect.Pointer, reflect.Interface:
		return v.IsNil()

	case reflect.Bool:
		return !v.Bool()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0

	case reflect.Float32, reflect.Float64:
		return v.Float() == 0

	case reflect.Struct:
		zero := reflect.Zero(v.Type())
		return reflect.DeepEqual(v.Interface(), zero.Interface())

	}

	return false
}

func RemoveExtraSpaces(s string) string {
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")
	return strings.TrimSpace(s)
}
