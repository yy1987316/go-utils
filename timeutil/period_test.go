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
		want PeriodTs
	}{
		{"Case1", time1, PeriodTs{
			time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local).Unix(),
			time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local).Unix()},
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
		want map[int]PeriodTs
	}{
		{"Case1", time1, map[int]PeriodTs{
			1: {time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local).Unix(),
				time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local).Unix()},
			2: {time.Date(2020, 2, 1, 0, 0, 0, 0, time.Local).Unix(),
				time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local).Unix()},
			3: {time.Date(2020, 3, 1, 0, 0, 0, 0, time.Local).Unix(),
				time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local).Unix()},
			4: {time.Date(2020, 4, 1, 0, 0, 0, 0, time.Local).Unix(),
				time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local).Unix()},
			5: {time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local).Unix(),
				time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local).Unix()},
			6: {time.Date(2020, 6, 1, 0, 0, 0, 0, time.Local).Unix(),
				time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local).Unix()},
			7: {time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local).Unix(),
				time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local).Unix()},
			8: {time.Date(2020, 8, 1, 0, 0, 0, 0, time.Local).Unix(),
				time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local).Unix()},
			9: {time.Date(2020, 9, 1, 0, 0, 0, 0, time.Local).Unix(),
				time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local).Unix()},
			10: {time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local).Unix(),
				time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local).Unix()},
			11: {time.Date(2020, 11, 1, 0, 0, 0, 0, time.Local).Unix(),
				time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local).Unix()},
			12: {time.Date(2020, 12, 1, 0, 0, 0, 0, time.Local).Unix(),
				time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local).Unix()},
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
