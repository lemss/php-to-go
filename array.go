// IsArray 实现类似PHP的is_array()函数
func IsArray(v interface{}) bool {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	// case reflect.Map:
	// 	return true
	case reflect.Slice:
		return true
	default:

	}

	return false
}

// ArrayKeys 实现类似PHP的array_keys()函数
func ArrayKeys(array map[string]interface{}) []string {
	var keys []string
	for k := range array {
		keys = append(keys, k)
	}
	return keys
}
