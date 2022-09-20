// Time 实现类似PHP的time()函数
func Time() int64 {
	return time.Now().Unix()
}

// Strtotime 实现类似PHP的strtotime()函数
// Strtotime("02/01/2006 15:04:05", "02/01/2016 15:04:05") == 1451747045
// Strtotime("3 04 PM", "8 41 PM") == -62167144740
func Strtotime(format, strtime string) (int64, error) {
	t, err := time.Parse(format, strtime)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
