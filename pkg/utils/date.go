package utils

import "time"

// StandardTime 获取冬令时起止时间
func StandardTime(year int) (time.Time, time.Time) {
	var stDateStart, stDateEnd time.Time
	firstDayOfNovember := time.Date(year, time.November, 1, 2, 0, 0, 0, time.Local)
	if firstDayOfNovember.Weekday() == time.Sunday {
		stDateStart = firstDayOfNovember.AddDate(0, 0, 7)
	} else {
		stDateStart = firstDayOfNovember.AddDate(0, 0, int(7-firstDayOfNovember.Weekday()))
	}
	firstDayOfMarch := time.Date(year+1, time.March, 1, 2, 0, 0, 0, time.Local)
	if firstDayOfMarch.Weekday() == time.Sunday {
		stDateEnd = firstDayOfMarch.AddDate(0, 0, 7)
	} else {
		stDateEnd = firstDayOfMarch.AddDate(0, 0, 7+int(7-firstDayOfMarch.Weekday()))
	}
	return stDateStart, stDateEnd
}
