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
