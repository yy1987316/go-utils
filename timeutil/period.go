package timeutil

import "time"

type PeriodTs struct {
	StartTime int64
	EndTime   int64
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
func GetMonthTs(t time.Time) PeriodTs {
	year := t.Year()
	month := t.Month()
	startTime := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	endTime := time.Date(year, month+1, 1, 0, 0, 0, 0, time.Local)
	return PeriodTs{startTime.Unix(), endTime.Unix()}
}

// GetYearTs returns a period map of every month of the year of t, with key in [1,12]
func GetYearTs(t time.Time) map[int]PeriodTs {
	tsMap := make(map[int]PeriodTs)
	year := t.Year()
	for i := 1; i < 13; i++ {
		startTime := time.Date(year, time.Month(i), 1, 0, 0, 0, 0, time.Local)
		endTime := time.Date(year, time.Month(i)+1, 1, 0, 0, 0, 0, time.Local)
		tsMap[i] = PeriodTs{startTime.Unix(), endTime.Unix()}
	}
	return tsMap
}
