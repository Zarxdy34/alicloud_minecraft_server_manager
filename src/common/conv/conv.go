package conv

import "strconv"

func StrToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func Int64ToStr(i int64) string {
	s := strconv.FormatInt(i, 10)
	return s
}
