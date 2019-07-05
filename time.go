package gotime

import "time"

type gotime struct {
	time.Time
}

// Now returns the current local time.
func Now() *gotime {
	return &gotime{time.Now()}
}

func (t *gotime) JST() time.Time {
	// TODO JST
	jst := time.Location{}
	return t.Time.In(&jst)
}
