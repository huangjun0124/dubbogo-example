package util

import "time"

func GetCommonNowStr() string{
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetCommonTimeStr(t time.Time) string{
	return t.Format("2006-01-02 15:04:05")
}
