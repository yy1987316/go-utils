package timeutil

// GetDayCount returns day count of the month, or of the year
// if input month == 0
func GetDayCount(year, month int) int {
	if month < 0 || month > 12 {
		return 0
	}

	switch month {
	case 0:
		if IsLeapYear(year) {
			return 366
		} else {
			return 365
		}
	case 2:
		if IsLeapYear(year) {
			return 29
		} else {
			return 28
		}
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	default:
		return 30
	}
}

// IsLeapYear returns the input year is whether a leap year or not
func IsLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}
