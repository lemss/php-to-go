// Gethostname 实现类似PHP的gethostname()函数
func Gethostname() (string, error) {
	return os.Hostname()
}

// Gethostbyname 实现类似PHP的gethostbyname()函数
// 获取与给定 Internet 主机名对应的 IPv4 地址
func Gethostbyname(hostname string) (string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		for _, v := range ips {
			if v.To4() != nil {
				return v.String(), nil
			}
		}
		return "", nil
	}
	return "", err
}
