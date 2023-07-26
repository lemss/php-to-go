// Echo 实现类似PHP的echo函数
func Echo(args ...interface{}) {
	fmt.Print(args...)
}

// Uniqid 实现类似PHP的uniqid函数
func Uniqid(prefix string) string {
	now := time.Now()
	return fmt.Sprintf("%s%08x%05x", prefix, now.Unix(), now.UnixNano()%0x100000)
}

// Exit 实现类似PHP的exit函数
func Exit(status int) {
	os.Exit(status)
}

// Die 实现类似PHP的die函数
func Die(status int) {
	os.Exit(status)
}

// Getenv 实现类似PHP的getenv函数
func Getenv(varname string) string {
	return os.Getenv(varname)
}
