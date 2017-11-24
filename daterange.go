package utils

import "time"

// DateRange returns a range of dates
func DateRange(startDate time.Time, endDate time.Time) []string {
	dates := []string{startDate.UTC().Format("2006-01-02")}
	c := startDate
	end := endDate.UTC().Format("2006-01-02")

	for c.UTC().Format("2006-01-02") != end {
		c = c.AddDate(0, 0, 1)
		dates = append(dates, c.UTC().Format("2006-01-02"))
	}

	return dates
}
