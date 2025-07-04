package db

import (
	"fmt"
	"gorm/types"
	"regexp"
	"unicode/utf8"
)

// Validators are different from Constraints. Constraints are
// applied at the DB-level and instantiated at table creation.
// Validators are custom functions executed when you create a
// new object.

func ValidateObject[T types.Struct](table types.Table[T], obj *T) error {
	v := getReflectValue(obj)

	for _, col := range table.Cols {
		field := v.FieldByName(col.FieldName)
		value := field.String()

		if col.Validators.IsEmail {
			if !validateEmail(value) {
				return fmt.Errorf("email is invalid")
			}
		}
		if col.Validators.IsURL {
			if !validateURL(value) {
				return fmt.Errorf("url is invalid")
			}
		}

		if col.Validators.MinLength > 0 {
			if utf8.RuneCountInString(value) < col.Validators.MinLength {
				return fmt.Errorf("minimum length should be: %d", col.Validators.MinLength)
			}
		}

		if col.Validators.MaxLength > 0 {
			if utf8.RuneCountInString(value) > col.Validators.MaxLength {
				return fmt.Errorf("maximum length should be: %d", col.Validators.MaxLength)
			}
		}

	}

	return nil
}

func validateEmail(s string) bool {
	re := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return re.MatchString(s)
}

func validateURL(s string) bool {
	re := regexp.MustCompile(`^https?:\/\/[a-z0-9]+(?:[-.][a-z0-9]+)*(?::[0-9]{1,5})?(?:\/[^\/\r\n]+)*\.[a-z]{2,5}(?:[?#]\S*)?$`)
	return re.MatchString(s)
}
