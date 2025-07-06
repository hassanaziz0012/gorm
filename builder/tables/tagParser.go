package tables

import (
	"fmt"
	"gorm/db"
	"gorm/types"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func parseCols(t reflect.Type) []types.Column {
	var cols []types.Column

	for i := 0; i < t.NumField(); i++ {
		var col types.Column
		if colIsForeignKey(t, i) {
			col = parseForeignKey(t, i)
		} else {
			col = parseCol(t, i)
		}
		cols = append(cols, col)
	}
	return cols
}

func parseCol(t reflect.Type, i int) types.Column {
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
	return col
}

func parseForeignKey(t reflect.Type, i int) types.Column {
	field := t.Field(i)
	colname := db.ToSnakeCase(field.Name) + "_id"
	coltype := types.Integer

	constraints := parseConstraints(field.Tag)
	validators := parseValidators(field.Tag)
	defaults := parseDefaults(field.Tag)

	fkTableName := GenerateTableName(field.Type)
	foreignkey := types.FKR{
		ConstraintName: "fk_" + fkTableName,
		Name:           colname,
		FKTable:        fkTableName,
		FKField:        "id",
		OnDelete:       types.CASCADE,
	}

	col := types.Column{
		FieldName:   field.Name,
		Name:        colname,
		Coltype:     coltype,
		Constraints: constraints,
		Defaults:    defaults,
		Validators:  validators,
		FKR:         foreignkey,
	}

	return col
}

func colIsForeignKey(t reflect.Type, i int) bool {
	field := t.Field(i)
	_, err := db.ParseDataType(field.Type.Name())
	if err != nil {
		fmt.Println(field.Type)
		fkTable := GenerateTableName(field.Type)
		if Exists(fkTable) {
			return true

		} else {
			log.Fatal("invalid data type: ", field.Type.Name())
		}
	}
	return false
}

func parseConstraints(tag reflect.StructTag) types.Constraints {
	c := types.Constraints{}
	c.Required = true // all fields required by default.

	for t := range strings.SplitSeq(tag.Get("gorm.constraints"), ",") {
		if t == "nullable" {
			c.Required = false
		}
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
		d.Bool = types.TRUE
	case "false":
		d.Bool = types.FALSE
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
