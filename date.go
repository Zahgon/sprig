package sprig

import (
	"time"
)

// Given a format and a date, format the date string.
//
// Date can be a `time.Time` or an `int, int32, int64`.
// In the later case, it is treated as seconds since UNIX
// epoch.
func date(fmt string, date interface{}) string { _ = "STUB: not implemented"; return "" }

func htmlDate(date interface{}) string { _ = "STUB: not implemented"; return "" }

func htmlDateInZone(date interface{}, zone string) string { _ = "STUB: not implemented"; return "" }

func dateInZone(fmt string, date interface{}, zone string) string {
	_ = "STUB: not implemented"
	return ""
}

func dateModify(fmt string, date time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func mustDateModify(fmt string, date time.Time) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func dateAgo(date interface{}) string { _ = "STUB: not implemented"; return "" }

// Drop resolution to seconds

func duration(sec interface{}) string { _ = "STUB: not implemented"; return "" }

func durationRound(duration interface{}) string { _ = "STUB: not implemented"; return "" }

func toDate(fmt, str string) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func mustToDate(fmt, str string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func unixEpoch(date time.Time) string { _ = "STUB: not implemented"; return "" }
