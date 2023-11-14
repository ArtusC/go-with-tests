package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	v := reflect.ValueOf(x)

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fn(f.String())
	}
}
