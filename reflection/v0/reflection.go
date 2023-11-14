package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	v := reflect.ValueOf(x)
	f := v.Field(0)
	fn(f.String())
}
