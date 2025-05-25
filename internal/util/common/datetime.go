package common

import (
	"learn-fiber/pkg/logger"
	"time"
)

func IsBeforeNow(t time.Time) bool {
	utcNow := time.Now().In(time.UTC)
	return t.Before(utcNow)
}

func IsAfterNow(t time.Time) bool {
	utcNow := time.Now().In(time.UTC)
	return t.After(utcNow)
}

func IsBefore(t1, t2 time.Time) bool {
	return t1.Before(t2)
}

func IsAfter(t1, t2 time.Time) bool {
	return t1.After(t2)
}

// Check datetime is before current date
// @params timeStr string of datetime
// @return boolean, err
func IsBeforeNowStr(timeStr string) (result bool, err error) {
	t, err := parseIso8601Time(timeStr)
	if err != nil {
		logger.L.Error(err)
		return false, err
	}
	return IsBeforeNow(t), nil
}

func IsAfterNowStr(timeStr string) (result bool, err error) {
	t, err := parseIso8601Time(timeStr)
	if err != nil {
		logger.L.Error(err)
		return false, err
	}
	return IsAfterNow(t), nil
}

func IsBeforeStr(tStr1, tStr2 string) (result bool, err error) {
	t1, err := parseIso8601Time(tStr1)
	if err != nil {
		logger.L.Error(err)
		return false, err
	}
	t2, err := parseIso8601Time(tStr2)
	if err != nil {
		logger.L.Error(err)
		return false, err
	}
	return IsBefore(t1, t2), nil
}

func IsAfterStr(tStr1, tStr2 string) (result bool, err error) {
	t1, err := parseIso8601Time(tStr1)
	if err != nil {
		logger.L.Error(err)
		return false, err
	}
	t2, err := parseIso8601Time(tStr2)
	if err != nil {
		logger.L.Error(err)
		return false, err
	}
	return IsAfter(t1, t2), nil
}

func parseIso8601Time(value string) (time.Time, error) {
	return time.Parse(time.RFC3339, value)
}
