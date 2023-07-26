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

// ArrayFill 实现类似PHP的array_fill()函数
func ArrayFill(startIndex int, num uint, value interface{}) map[int]interface{} {
	m := make(map[int]interface{})
	var i uint
	for i = 0; i < num; i++ {
		m[startIndex] = value
		startIndex++
	}
	return m
}

// ArrayFlip 实现类似PHP的array_flip()函数
func ArrayFlip(m map[interface{}]interface{}) map[interface{}]interface{} {
	n := make(map[interface{}]interface{})
	for i, v := range m {
		n[v] = i
	}
	return n
}
