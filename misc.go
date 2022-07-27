// Echo 实现类似PHP的echo函数
func Echo(args ...interface{}) {
	fmt.Print(args...)
}

// Uniqid 实现类似PHP的uniqid函数
func Uniqid(prefix string) string {
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", prefix, now.Unix(), now.UnixNano()%0x100000)
}
