// Serialize 实现类似PHP的serialize()函数
func Serialize(value interface{}) ([]byte, error) {
	return serialize.Marshal(value)
}

// Unserialize 实现类似PHP的unserialize()函数
func Unserialize(data []byte) (interface{}, error) {
	return serialize.UnMarshal(data)
}
