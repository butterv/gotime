package gotime

import "time"

type gotime struct {
	t time.Time
}

// Now returns the current local time.
func Now() *gotime {
	return &gotime{t: time.Now()}
}

func Date(year, month, day, hour, min, sec uint) *gotime {
	location := time.Local
	date := time.Date(int(year), time.Month(int(month)), int(day), int(hour), int(min), int(sec), 0, location)
	return &gotime{t: date}
}

func (g *gotime) In(zone string) (time.Time, error) {
	location, err := time.LoadLocation(zone)
	if err != nil {
		return time.Time{}, err
	}
	return g.t.In(location), nil
}

func (g *gotime) AddYear(years uint) *gotime {
	date := time.Date(g.t.Year()+int(years), g.t.Month(), g.t.Day(), g.t.Hour(), g.t.Minute(), g.t.Second(), g.t.Nanosecond(), g.t.Location())
	return &gotime{t: date}
}

func (g *gotime) AddMonth(months uint) *gotime {
	date := time.Date(g.t.Year(), time.Month(int(g.t.Month())+int(months)), g.t.Day(), g.t.Hour(), g.t.Minute(), g.t.Second(), g.t.Nanosecond(), g.t.Location())
	return &gotime{t: date}
}

func (g *gotime) AddDay(days uint) *gotime {
	date := time.Date(g.t.Year(), g.t.Month(), g.t.Day()+int(days), g.t.Hour(), g.t.Minute(), g.t.Second(), g.t.Nanosecond(), g.t.Location())
	return &gotime{t: date}
}
