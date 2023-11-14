package reflection

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {

	val := extractValue(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

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

func extractValue(valInput interface{}) reflect.Value {
	valOut := reflect.ValueOf(valInput)

	if valOut.Kind() == reflect.Pointer {
		valOut = valOut.Elem()
	}

	return valOut
}
