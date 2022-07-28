// Serialize 实现类似PHP的serialize()函数
func Serialize(value interface{}) ([]byte, error) {
	return serialize.Marshal(value)
}
