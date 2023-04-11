package helper

import (
	"fmt"
	"strings"
	"time"
)

var IndonesiaWeekday = map[time.Weekday]string{
	time.Monday:    "Senin",
	time.Tuesday:   "Selasa",
	time.Wednesday: "Rabu",
	time.Thursday:  "Kamis",
	time.Friday:    "Jumat",
	time.Saturday:  "Sabtu",
	time.Sunday:    "Minggu",
}

var IndonesiaMonth = map[time.Month]string{
	time.January:   "Januari",
	time.February:  "Februari",
	time.March:     "Maret",
	time.April:     "April",
	time.May:       "Mei",
	time.June:      "Juni",
	time.July:      "Juli",
	time.August:    "Agustus",
	time.September: "September",
	time.October:   "Oktober",
	time.November:  "November",
	time.December:  "Desember",
}

func TimeToIndonesiaLang(t time.Time, layout string) string {
	d := t.Format(layout)

	d = strings.ReplaceAll(d, t.Month().String(), IndonesiaMonth[t.Month()])
	d = strings.ReplaceAll(d, t.Weekday().String(), IndonesiaWeekday[t.Weekday()])

	return d
}

func TimeToFormatted(t time.Time) string {
	year, month, day := t.Date()

	if t.IsZero() {
		return ""
	}
	return fmt.Sprintf("%d %s %d", day, IndonesiaMonth[month], year)
}

func TimeToFormatted2(t time.Time) string {
	year, month, day := t.Date()

	if t.IsZero() {
		return ""
	}
	return fmt.Sprintf("%d %s %d | %s WIB", day, IndonesiaMonth[month], year, t.Format("15:04"))
}

func TimeToFormatted3(t time.Time) string {
	year, month, day := t.Date()

	if t.IsZero() {
		return ""
	}
	return fmt.Sprintf("%d %s %d, %s", day, IndonesiaMonth[month], year, t.Format("15:04"))
}
