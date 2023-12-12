package reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	v := reflect.ValueOf(x)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}
	}
}
