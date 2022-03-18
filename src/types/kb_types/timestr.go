package kb_types

import "time"

type TimeStr string

func newTimeStr(t time.Time) TimeStr {
	return TimeStr(t.Format("2006-01-02T15:04:05-07:00"))
}

func (t TimeStr) String() string {
	return string(t)
}
