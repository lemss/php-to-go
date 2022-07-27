// Gethostname 实现类似PHP的gethostname()函数
func Gethostname() (string, error) {
	return os.Hostname()
}
