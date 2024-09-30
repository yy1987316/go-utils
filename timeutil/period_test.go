package timeutil

import (
	"reflect"
	"testing"
	"time"
)

func TestPeriodTs_Cover(t *testing.T) {
	type fields struct {
		StartTime int64
		EndTime   int64
	}
	type args struct {
		p2 *PeriodTs
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PeriodTs
	}{
		{"Case1", fields{20, 100}, args{&PeriodTs{200, 500}}, nil},
		{"Case2", fields{20, 100}, args{&PeriodTs{30, 110}}, &PeriodTs{30, 100}},
		{"Case3", fields{20, 100}, args{&PeriodTs{10, 40}}, &PeriodTs{20, 40}},
		{"Case4", fields{20, 100}, args{&PeriodTs{50, 60}}, &PeriodTs{50, 60}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PeriodTs{
				StartTime: tt.fields.StartTime,
				EndTime:   tt.fields.EndTime,
			}
			if got := p.Cover(tt.args.p2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cover() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMonthTs(t *testing.T) {
	time1, _ := time.Parse("2006-01-02 15:04:05", "2020-12-09 00:32:00")
	tests := []struct {
		name string
		args time.Time
		want MonthPeriodTs
	}{
		{"Case1", time1, MonthPeriodTs{
			2020,
			12,
			PeriodTs{
				time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local).Unix(),
				time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local).Unix()},
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMonthTs(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMonthTs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYearTs(t *testing.T) {
	time1, _ := time.Parse("2006-01-02 15:04:05", "2020-12-09 00:32:00")
	tests := []struct {
		name string
		args time.Time
		want []MonthPeriodTs
	}{
		{"Case1", time1, []MonthPeriodTs{
			{2020, 1,
				PeriodTs{time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
					time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local).Unix()}},
			{2020, 2,
				PeriodTs{time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local).Unix(),
					time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local).Unix()}},
			{2020, 3,
				PeriodTs{time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local).Unix(),
					time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local).Unix()}},
			{2020, 4,
				PeriodTs{time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local).Unix(),
					time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local).Unix()}},
			{2020, 5,
				PeriodTs{time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local).Unix(),
					time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local).Unix()}},
			{2020, 6,
				PeriodTs{time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local).Unix(),
					time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local).Unix()}},
			{2020, 7,
				PeriodTs{time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local).Unix(),
					time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local).Unix()}},
			{2020, 8,
				PeriodTs{time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local).Unix(),
					time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local).Unix()}},
			{2020, 9,
				PeriodTs{time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local).Unix(),
					time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local).Unix()}},
			{2020, 10,
				PeriodTs{time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local).Unix(),
					time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local).Unix()}},
			{2020, 11,
				PeriodTs{time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local).Unix(),
					time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local).Unix()}},
			{2020, 12,
				PeriodTs{time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local).Unix(),
					time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local).Unix()}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetYearTs(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetYearTs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLastNMonthTs(t *testing.T) {
	time1, _ := time.Parse("2006-01-02 15:04:05", "2020-12-09 00:32:00")
	time2, _ := time.Parse("2006-01-02 15:04:05", "2021-01-30 00:32:00")

	type args struct {
		n int
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want []MonthPeriodTs
	}{
		{
			"Case1", args{12, time1}, GetYearTs(time1),
		},
		{
			"Case2", args{12, time2}, []MonthPeriodTs{
				{2020, 2,
					PeriodTs{time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local).Unix(),
						time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local).Unix()}},
				{2020, 3,
					PeriodTs{time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local).Unix(),
						time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local).Unix()}},
				{2020, 4,
					PeriodTs{time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local).Unix(),
						time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local).Unix()}},
				{2020, 5,
					PeriodTs{time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local).Unix(),
						time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local).Unix()}},
				{2020, 6,
					PeriodTs{time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local).Unix(),
						time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local).Unix()}},
				{2020, 7,
					PeriodTs{time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local).Unix(),
						time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local).Unix()}},
				{2020, 8,
					PeriodTs{time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local).Unix(),
						time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local).Unix()}},
				{2020, 9,
					PeriodTs{time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local).Unix(),
						time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local).Unix()}},
				{2020, 10,
					PeriodTs{time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local).Unix(),
						time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local).Unix()}},
				{2020, 11,
					PeriodTs{time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local).Unix(),
						time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local).Unix()}},
				{2020, 12,
					PeriodTs{time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local).Unix(),
						time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local).Unix()}},
				{2021, 1,
					PeriodTs{time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
						time.Date(2021, 2, 1, 0, 0, 0, 0, time.Local).Unix()}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLastNMonthTs(tt.args.n, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLastNMonthTs() = %v, want %v", got, tt.want)
			}
		})
	}
}
