package mappy

import (
	stderr "errors"
	"reflect"

	"github.com/pkg/errors"
)

// ErrMapMarshal is returned when it is not possible to marshal struct into a map.
var ErrMapMarshal = stderr.New("failed to marshal struct into map")

// Marshal transforms a custom struct into string->string map.
func Marshal(data interface{}) (m map[string]string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.Wrapf(ErrMapMarshal, "%v", r)
		}
	}()

	m = make(map[string]string)
	rv := reflect.ValueOf(data)
	elT := reflect.TypeOf(data)
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		ftype := elT.Field(i)

		tag := ftype.Tag.Get("map")
		if tag != "" {
			m[tag] = field.String()
		}
	}
	return m, nil
}
