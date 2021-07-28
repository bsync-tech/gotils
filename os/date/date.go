package date

import (
	"time"

	"github.com/araddon/dateparse"
)

var TIME_FORMAT_SECOND = "2006-01-02 15:34:05.000"
var TIME_FORMAT_DATE = "2006-01-02"

func SecondsNow() *time.Time {
	time.LoadLocation("Asia/Shanghai")
	t, err := time.Parse(TIME_FORMAT_SECOND, time.Now().String())
	if nil != err {
		t = time.Time{}
		return &t
	}

	return &t
}

func TodayWithFormat(format string) *time.Time {
	t, err := time.Parse(format, time.Now().Format(format))
	if nil != err {
		t = time.Time{}
		return &t
	}

	return &t
}

var DEFAULT_LOCATION = "Asia/Shanghai"

func Parse(d string) (*time.Time, error) {
	loc, err := time.LoadLocation(DEFAULT_LOCATION)
	if err != nil {
		return nil, err
	}
	time.Local = loc
	t, err := dateparse.ParseLocal(d)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func Today() string {
	loc, err := time.LoadLocation(DEFAULT_LOCATION)
	if err != nil {
		return ""
	}
	time.Local = loc
	return time.Now().Format(TIME_FORMAT_DATE)
}
