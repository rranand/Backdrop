package formatter

import "time"

func dateTimeFormatter(dateTime time.Time) string {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	istTime := dateTime.In(loc)

	return istTime.Format("2006-01-02 15:04:05")
}
