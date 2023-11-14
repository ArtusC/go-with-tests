package reflection

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	v := reflect.ValueOf(x)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			fn(fmt.Sprint(field.Int()))
		case reflect.Float32, reflect.Float64:
			fn(fmt.Sprintf("%.2f", field.Float()))
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
