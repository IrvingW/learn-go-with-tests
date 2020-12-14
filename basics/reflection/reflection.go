package reflection

import "reflect"

// Walk will iterate every field in struct x, and call fn on that field
// 使用interface{}指代可以传入任何东西
func Walk(x interface{}, fn func(string)) {
	value := reflect.ValueOf(x)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	number := 0
	var getField func(int) reflect.Value
	switch value.Kind() {
	case reflect.String:
		fn(value.String())
	case reflect.Struct:
		number = value.NumField()
		getField = value.Field
	case reflect.Array, reflect.Slice:
		number = value.Len()
		getField = value.Index
	case reflect.Map:
		for _, key := range value.MapKeys() {
			Walk(value.MapIndex(key).Interface(), fn)
		}
	}

	for i := 0; i < number; i++ {
		Walk(getField(i).Interface(), fn)
	}
}
