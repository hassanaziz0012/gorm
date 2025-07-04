package tables

import (
	"gorm/db"
	types "gorm/types"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func parseCols(t reflect.Type) []types.Column {
	var cols []types.Column

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		fieldName := field.Name
		colname := db.ToSnakeCase(field.Name)
		coltype, err := db.ParseDataType(field.Type.Name())
		if err != nil {
			log.Fatal("invalid data type: ", field.Type.Name())
		}
		constraints := parseConstraints(field.Tag)
		validators := parseValidators(field.Tag)
		defaults := parseDefaults(field.Tag)

		col := types.Column{
			FieldName:   fieldName,
			Name:        colname,
			Coltype:     coltype,
			Constraints: constraints,
			Defaults:    defaults,
			Validators:  validators,
		}
		cols = append(cols, col)
	}
	return cols
}

func parseConstraints(tag reflect.StructTag) types.Constraints {
	c := types.Constraints{}
	for t := range strings.SplitSeq(tag.Get("gorm.constraints"), ",") {
		if t == "pk" {
			c.IsPrimary = true
		}
		if t == "autoincrement" {
			c.AutoIncrement = true
		}
		if t == "unique" {
			c.IsUnique = true
		}
	}

	return c
}

func parseDefaults(tag reflect.StructTag) types.Defaults {
	d := types.Defaults{}
	t := tag.Get("gorm.default")

	switch t {
	case "now":
		d.TimeNow = true
	case "true":
		d.Bool = true
	case "false":
		d.Bool = false
	default:
		if i, err := strconv.Atoi(t); err == nil {
			d.Int = i
		} else {
			d.Text = t
		}
	}

	return d
}

func parseValidators(tag reflect.StructTag) types.Validators {
	var v types.Validators
	for val := range strings.SplitSeq(tag.Get("gorm.validators"), ",") {
		if val == "email" {
			v.IsEmail = true
		}
		if val == "url" {
			v.IsURL = true
		}
		if strings.HasPrefix(val, "min(") && strings.HasSuffix(val, ")") {
			extracted := strings.TrimSuffix(strings.TrimPrefix(val, "min("), ")")
			minLength, err := strconv.Atoi(extracted)
			if err != nil {
				log.Fatal("unable to parse min() value: ", err)
			}
			v.MinLength = minLength
		}
		if strings.HasPrefix(val, "max(") && strings.HasSuffix(val, ")") {
			extracted := strings.TrimSuffix(strings.TrimPrefix(val, "max("), ")")
			maxLength, err := strconv.Atoi(extracted)
			if err != nil {
				log.Fatal("unable to parse max() value: ", err)
			}
			v.MaxLength = maxLength
		}
	}
	return v
}
