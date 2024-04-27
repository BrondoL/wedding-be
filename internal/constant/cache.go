package constant

import "time"

const (
	CacheKeyAttendance = "attendance"
)

const (
	CacheTTLOneDay     = 24 * time.Hour
	CacheTTLFiveMinute = 5 * time.Minute
	CacheTTLForever    = 0
	CacheTTLInvalidate = -1
)
