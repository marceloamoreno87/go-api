package helper

import "strconv"

func StrToInt32(s string) (i int32) {
	i64, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}
	i = int32(i64)
	return
}
