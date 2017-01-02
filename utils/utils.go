package utils

import (
	"strconv"
	"time"
)

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func MergeMap(m1, m2 map[string]string) map[string]string {
	ans := map[string]string{}

	for k, v := range m1 {
		ans[k] = v
	}
	for k, v := range m2 {
		ans[k] = v
	}
	return (ans)
}

func MsToTime(ms int64) (time.Time, error) {
	str := Int64ToString(ms)
	msInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(0, msInt*int64(time.Millisecond)), nil
}
