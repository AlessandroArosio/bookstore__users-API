package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z" // just placeholders for time formatting
	apiDbLayout   = "2006-01-02 15:01:04"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDbFormat() string {
	return GetNow().Format(apiDbLayout)
}
