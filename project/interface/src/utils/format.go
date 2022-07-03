package utils

import (
	"strconv"
	"time"
)

func Format(v interface{}) time.Time {

	var t int
	switch v := v.(type) {
	case int:
		t = v
	case string:
		t, _ = strconv.Atoi(v)
	default:
		panic("type is unknown")
	}

	u := time.Unix(int64(t), 0)
	return u
}
