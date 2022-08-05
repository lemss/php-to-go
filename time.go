// Time 实现类似PHP的time()函数
func Time() int64 {
	return time.Now().Unix()
}
