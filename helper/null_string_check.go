package helper

import (
	"database/sql"
	"reflect"
	"strings"
)

func NullStringCheck[T any](any T) {
	object := reflect.ValueOf(any).Elem()

	for i := 0; i < object.NumField(); i++ {
		field := object.Field(i)
		if field.Type() == reflect.TypeOf(sql.NullString{}) && len(strings.TrimSpace(field.Interface().(sql.NullString).String)) == 0 {
			field.Set(reflect.ValueOf(sql.NullString{Valid: false}))
		}
	}
}
