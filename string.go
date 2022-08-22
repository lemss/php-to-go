// Serialize 实现类似PHP的serialize()函数
func Serialize(value interface{}) ([]byte, error) {
	return serialize.Marshal(value)
}

// Unserialize 实现类似PHP的unserialize()函数
func Unserialize(data []byte) (interface{}, error) {
	return serialize.UnMarshal(data)
}

// Strpos 实现类似PHP的strpos()函数
func Strpos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}
