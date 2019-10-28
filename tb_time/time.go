package tb_time

import (
	"time"
)

const TimeStyles string = "2006-01-02 15:04:05"
const ZEROTimeStr = "0000-00-00 00:00:00"
const NULLTimeStr = "0001-01-01 00:00:00"
const DBDefaultTimeStr string = "2006-01-02 15:04:05"

// TimeNowCST return now beijing time
func TimeNowCST() (time.Time, error) {
	var curNow time.Time
	timelocal, err := time.LoadLocation("Asia/Chongqing")
	if err != nil {
		return curNow, err
	}
	time.Local = timelocal
	curNow = time.Now().Local()
	return curNow, nil
}

// TransToCST trans to beijing time
// cstStr is beijing time string like: 2017-06-20 18:16:15
func TransToCST(cstStr string) (time.Time, error) {
	if cstStr == "" {
		return ZEROTime(), nil
	}
	t := time.Time{}
	local := time.UTC
	tmp, err := time.ParseInLocation(TimeStyles, cstStr, local)
	if err != nil {
		return t, err
	}
	t = tmp
	return t, nil
}

// CSTTransToA trans to CST to string
// cst is a beijing time
func CSTTransToA(cst time.Time) string {
	return cst.Format(TimeStyles)
}

func ZEROTime() time.Time {
	z, _ := TransToCST(ZEROTimeStr)
	return z
}

func IsNULLTime(t time.Time) bool {
	if t.Format(TimeStyles) == NULLTimeStr {
		return true
	} else {
		return false
	}
}

func NewDBDefaultTime() time.Time {
	z, _ := TransToCST(TimeStyles)
	return z
}

func IsDBDefaultTime(t time.Time) bool {
	if t.Format(TimeStyles) == TimeStyles {
		return true
	} else {
		return false
	}
}

func DBDefaultTimeTransZEROTime(t time.Time) time.Time {
	if IsDBDefaultTime(t) {
		return ZEROTime()
	}
	return t
}

func ZEROTimeTransDBDefaultTime(t time.Time) time.Time {
	if IsNULLTime(t) {
		return NewDBDefaultTime()
	} else {
		return t
	}
}
