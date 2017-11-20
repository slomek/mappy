package mappy

import "reflect"

// Marshal transforms a custom struct into string->string map.
func Marshal(data interface{}) (map[string]string, error) {
	m := map[string]string{}
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
