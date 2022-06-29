package strconve

import "strconv"

func StringToInt64(s string) (int64, error) {
	num, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}
	return num, err
}
