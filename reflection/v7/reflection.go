package reflection

import (
	"fmt"
	"reflect"
)

func walk(x interface{}, fn func(input string)) {

	val := extractValue(x)

	walkValue := func(v reflect.Value) {
		walk(v.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		fn(fmt.Sprint(val.Int()))
	case reflect.Float32, reflect.Float64:
		fn(fmt.Sprintf("%.2f", val.Float()))
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
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
