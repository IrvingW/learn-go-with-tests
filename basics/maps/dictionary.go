package maps

// Dict 自定义类型
type Dict map[string]string

// DictError 自定义错误类型
type DictError string

func (e DictError) Error() string {
	return string(e)
}

const (
	// ErrorNotFound is returned when could not search out any result using given key
	ErrorNotFound = DictError("could not find the key you are looking for")

	// ErrorKeyExist is returned when insert key already exists in map
	ErrorKeyExist = DictError("the key you are inserting already exist")
)

// Search accept a string argument, then using this input search in map
// return the value if searching hit
// return "" and ErrorNotFound error if could not hit any record in map
// map非常特殊，我不需要在指针上调用他的方法，因为map里面的值都是引用类型
// quote "Map 有一个有趣的特性，不使用指针传递你就可以修改它们。这是因为 map 是引用类型。这意味着它拥有对底层数据结构的引用，就像指针一样。"
func (d Dict) Search(key string) (string, error) {
	result, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}
	return result, nil
}

// Add insert a new pair of string into map
func (d Dict) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrorNotFound:
		d[key] = value
	case nil:
		return ErrorKeyExist
	default:
		return err
	}
	return nil
}
