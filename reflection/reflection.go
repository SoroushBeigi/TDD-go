package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(i interface{}) reflect.Value {
	val := reflect.ValueOf(i)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
