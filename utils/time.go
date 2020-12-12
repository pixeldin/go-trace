package utils

import "time"

// beginTime till now's duration
func DiffNano(beginTime time.Time) int64 {
	return int64(time.Since(beginTime))
}
