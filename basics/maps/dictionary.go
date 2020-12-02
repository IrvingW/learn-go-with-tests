package maps

import (
	"errors"
)

// Dict 自定义类型
type Dict map[string]string

// ErrorNotFound is returned when could not search out any result using given key
var ErrorNotFound = errors.New("could not find the key you are looking for.")

// Search accept a string argument, then using this input search in map
// return the value if searching hit
// return "" and ErrorNotFound error if could not hit any record in map
// map非常特殊，我不需要在指针上调用他的方法，因为map里面的值都是引用类型
func (d Dict) Search(key string) (string, error) {
	result, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}
	return result, nil
}
