package mappy

import (
	"reflect"
)

// Unmarshal transforms a string->string map into a custom struct.
func Unmarshal(m map[string]string, data interface{}) error {
	rv := reflect.ValueOf(data)
	el := rv.Elem()
	elT := el.Type()
	for i := 0; i < el.NumField(); i++ {
		field := el.Field(i)
		ftype := elT.Field(i)

		tag := ftype.Tag.Get("map")
		val := m[tag]

		field.SetString(val)
	}
	return nil
}
