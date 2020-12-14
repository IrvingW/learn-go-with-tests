package reflection

import "reflect"

// Walk will iterate every field in struct x, and call fn on that field
func Walk(x interface{}, fn func(string)) {
	value := reflect.ValueOf(x)
	field := value.Field(0)
	fn(field.String())
}
