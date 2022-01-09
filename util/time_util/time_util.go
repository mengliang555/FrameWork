package time_util

import (
	"time"
)

func GetCurrentTimeToSecond() string {
	return time.Now().Format(DatetimeLayout)
}

func GetCurrentTimeToRec3339() string {
	return time.Now().Format(DatetimeLayoutRfc3339)
}
