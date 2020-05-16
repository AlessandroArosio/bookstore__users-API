package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z" // just placeholders for time formatting
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
