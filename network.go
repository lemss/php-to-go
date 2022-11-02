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

// Gethostbynamel 实现类似PHP的gethostbynamel()函数
// 获取与给定 Internet 主机名对应的 IPv4 地址
func Gethostbynamel(hostname string) ([]string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		var ipstrs []string
		for _, v := range ips {
			if v.To4() != nil {
				ipstrs = append(ipstrs, v.String())
			}
		}
		return ipstrs, nil
	}
	return nil, err
}

// Gethostbyaddr 实现类似PHP的gethostbyaddr()函数
// Get the Internet host name corresponding to a given IP address
func Gethostbyaddr(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if names != nil {
		return strings.TrimRight(names[0], "."), nil
	}
	return "", err
}
