package reflection

import "reflect"

// Walk will iterate every field in struct x, and call fn on that field
// 使用interface{}指代可以传入任何东西
func Walk(x interface{}, fn func(string)) {
	value := reflect.ValueOf(x)
	for i := 0; i < value.NumField(); i++ {
		fn(value.Field(i).String())
	}
}
