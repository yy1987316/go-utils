package timeutil

import "time"

type PeriodTs struct {
	StartTime int64
	EndTime   int64
}

type MonthPeriodTs struct {
	Year  int
	Month int
	PeriodTs
}

// Duration returns the period duration
func (p *PeriodTs) Duration() int64 {
	return p.EndTime - p.StartTime
}

// Cover returns the intersection period of p and p2,
// returns nil when there is no intersection
func (p *PeriodTs) Cover(p2 *PeriodTs) *PeriodTs {
	if p.StartTime > p2.EndTime || p.EndTime < p2.StartTime {
		return nil
	}
	startTime := p.StartTime
	endTime := p.EndTime
	if startTime < p2.StartTime {
		startTime = p2.StartTime
	}
	if endTime > p2.EndTime {
		endTime = p2.EndTime
	}
	return &PeriodTs{startTime, endTime}
}

// GetMonthTs returns a period of the month of t
func GetMonthTs(t time.Time) MonthPeriodTs {
	year := t.Year()
	month := t.Month()
	startTime := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	endTime := time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local)
	return MonthPeriodTs{year, int(month), PeriodTs{startTime.Unix(), endTime.Unix()}}
}

// GetYearTs returns a period slice of every month of the year of t
func GetYearTs(t time.Time) []MonthPeriodTs {
	tsList := make([]MonthPeriodTs, 0, 12)
	for i := 1; i < 13; i++ {
		tsList = append(tsList, GetMonthTs(time.Date(t.Year(), time.Month(i), 1, 0, 0, 0, 0, time.Local)))
	}
	return tsList
}

// GetLastNMonthTs returns a period slice of every month of last n month(s) before t (include the month of t)
func GetLastNMonthTs(n int, t time.Time) []MonthPeriodTs {
	if n <= 0 {
		return nil
	}
	tsList := make([]MonthPeriodTs, 0, n)
	for i := n; i > 0; i-- {
		p := t.AddDate(0, -(i - 1), 0)
		tsList = append(tsList, GetMonthTs(p))
	}
	return tsList
}
