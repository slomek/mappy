package mappy

import (
	"errors"
	"fmt"
	"reflect"
)

// ErrMapMarshal is returned when it is not possible to unmarshal map into a struct.
var ErrMapUnmarshal = errors.New("failed to unmarshal map into struct")

// Unmarshal transforms a string->string map into a custom struct.
func Unmarshal(m map[string]string, data interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v: %w", r, ErrMapUnmarshal)
		}
	}()

	// Exit if the map is empty.
	if len(m) == 0 {
		return nil
	}

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

	return
}
