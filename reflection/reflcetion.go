package reflection

import (
	"reflect"
)

func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	var numberOfValues int
	var getValue func(i int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getValue = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getValue = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getValue(i).Interface(), fn)
	}
}

func getValue(x interface{}) (val reflect.Value) {
	val = reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return
}
