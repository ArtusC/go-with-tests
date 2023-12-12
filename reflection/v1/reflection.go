package reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	v := reflect.ValueOf(x)
	f := v.Field(0)
	fn(f.String())
}
