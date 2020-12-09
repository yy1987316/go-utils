package timeutil

import "testing"

func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		name string
		year int
		want bool
	}{
		{ "Case1", 2019, false},
		{ "Case2", 2200, false},
		{ "Case3", 2000, true},
		{ "Case4", 2020, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLeapYear(tt.year); got != tt.want {
				t.Errorf("IsLeapYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDayCount(t *testing.T) {
	type args struct {
		year  int
		month int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case1", args{2019, -1}, 0},
		{"Case2", args{2019, 2}, 28},
		{"Case3", args{2020, 2}, 29},
		{"Case4", args{2020, 0}, 366},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDayCount(tt.args.year, tt.args.month); got != tt.want {
				t.Errorf("GetDayCount() = %v, want %v", got, tt.want)
			}
		})
	}
}