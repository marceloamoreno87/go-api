package helper

import "time"

func GetTimestamp() int64 {
	return time.Now().Unix()
}

func CompareTimestamp(first int64, second int64) bool {
	return first > second
}
