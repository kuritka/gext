// Package formats and parese date .
package date

import (
	"time"

	"github.com/pkg/errors"
)

// ErrDatatimeMalformed raies when datetime not satisfy the following formats:
// "02.01.2006", "20060102" or "2006-01-02".
var (
	ErrDatatimeMalformed = errors.New("date not in proper format")
	ErrWrongTimeInterval = errors.New("wrong time interval")
)

// IsoDateFormatter converts the date in the YYYMMDD format required by cascade
// application.
func IsoDateFormatter(datetime string) (time.Time, error) {
	t, ok := findByIsoLayout(datetime)
	if !ok {
		return time.Time{}, ErrDatatimeMalformed
	}
	return t, nil
}

// ToISO86012004BasicString returns ISO8601:2004 basic date format YYYYMMDD representation
func ToISO86012004BasicString(time time.Time) string {
	return time.Format("20060102")
}

// ToISO8601DateString returns ISO8601 short date format YYYY-MM-DD representation
func ToISO8601DateString(time time.Time) string {
	return time.Format("2006-01-02")
}

// ToDEFormatDateString returns German locale date format DD.MM.YYYY representation
func ToDEFormatDateString(time time.Time) string {
	return time.Format("02.01.2006")
}

// ToDEFormatDateString returns German locale date format DD.MM.YYYY HH:MM:SS representation
func ToDEFormatDateTimeSecondsString(time time.Time) string {
	return time.Format("02.01.2006 15:04:02")
}

// ToDEFormatDateString returns German locale date format DD.MM.YYYY HH:MM representation
func ToDEFormatDateTimeString(time time.Time) string {
	return time.Format("02.01.2006 15:04")
}

// ToISO8601DateTimeString returns ISO8601 date format YYYY-MM-DDTHH:MM:SSZ representation
func ToISO8601DateTimeString(time time.Time) string {
	return time.Format("2006-01-02T15:04:05-0700")
}

// IsISODate returns true when datatime parsed successfully against isoLyaout
// formats, otherwise false.
func IsISODate(datetime string) bool {
	_, ok := findByIsoLayout(datetime)
	return ok
}

var isoLayouts = []string{"02.01.2006", "20060102", "2006-01-02", "2006-01-02 15:04", "2006-01-02 15:04:05", time.RFC3339}

func findByIsoLayout(datetime string) (time.Time, bool) {
	for _, layout := range isoLayouts {
		if t, err := time.Parse(layout, datetime); err == nil {
			return t, true
		}
	}
	return time.Time{}, false
}

// ISODateStringBeforeToday returns true if ISO8601 date string YYYY-MM-DD is before today
func ISODateStringBeforeToday(datetime string) (bool, error) {
	date, err := IsoDateFormatter(datetime)
	if err != nil {
		return false, err
	}
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, date.Location())
	if date.Before(today) {
		return true, nil
	}
	return false, nil
}

// GetDateTimeFromToForTonight parses time from string and returns from-to date time interval for tonight.
func GetDateTimeFromToForTonight(current time.Time, fromTimeString, toTimeString string) (time.Time, time.Time, error) {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return time.Time{}, time.Time{}, errors.Wrap(err, "failed to load location time zone")
	}
	var fromTime, toTime, fromToday, toToday time.Time
	if fromTime, err = time.ParseInLocation("15:04", fromTimeString, loc); err != nil {
		return time.Time{}, time.Time{}, errors.Wrap(err, "wrong `from` time format")
	}
	if toTime, err = time.ParseInLocation("15:04", toTimeString, loc); err != nil {
		return time.Time{}, time.Time{}, errors.Wrap(err, "wrong `to` time format")
	}
	// get current time in the given time zone
	now := current.In(loc)
	// get time of noon in the given time zone
	noon := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, loc)
	// init from-to for today
	fromToday = fromTime.AddDate(now.Year(), int(now.Month())-1, now.Day()-1)
	toToday = toTime.AddDate(now.Year(), int(now.Month())-1, now.Day()-1)
	// make the time interval from-to
	if now.Before(noon) {
		return fromToday.AddDate(0, 0, -1), toToday, nil
	}
	return fromToday, toToday.AddDate(0, 0, 1), nil
}

// GetDateTimeFromTo parses time from string and returns from-to date time interval for one day.
func GetDateTimeFromTo(current time.Time, fromTimeString, toTimeString string) (time.Time, time.Time, error) {
	var err error
	var fromTime, toTime, fromToday, toToday time.Time
	if fromTime, err = time.Parse("15:04", fromTimeString); err != nil {
		return time.Time{}, time.Time{}, errors.Wrap(err, "wrong `from` time format")
	}
	if toTime, err = time.Parse("15:04", toTimeString); err != nil {
		return time.Time{}, time.Time{}, errors.Wrap(err, "wrong `to` time format")
	}
	// get current time in the UTC time zone
	now := current.In(time.UTC)
	// init from-to for today
	fromToday = fromTime.AddDate(now.Year(), int(now.Month())-1, now.Day()-1)
	toToday = toTime.AddDate(now.Year(), int(now.Month())-1, now.Day()-1)
	// make the time interval from-to
	if toToday.Before(fromToday) {
		return time.Time{}, time.Time{}, ErrWrongTimeInterval
	}
	return fromToday, toToday, nil
}
