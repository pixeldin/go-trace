package utils

import "time"

// beginTime till now's duration
func DiffNano(beginTime time.Time) int64 {
	return int64(time.Since(beginTime))
}

func GetPresentFormat() string {
	time.LoadLocation("local")
	return time.Now().Format("2006-01-02 15:04:05.99")
}
