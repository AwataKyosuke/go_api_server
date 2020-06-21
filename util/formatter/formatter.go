package formatter

import (
	"strings"
	"time"
)

func Format(t time.Time) string {
	weekdayja := strings.NewReplacer(
		"Sun", "日",
		"Mon", "月",
		"Tue", "火",
		"Wed", "水",
		"Thu", "木",
		"Fri", "金",
		"Sat", "土",
	)
	return weekdayja.Replace(t.Format("2006年1月2日(Mon) 15時04分"))
}
