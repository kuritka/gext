package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsoDateFormatter(t *testing.T) {
	//ARRANGE
	cases := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{name: "Date format type 1", input: "02.01.2006", output: "20060102", err: nil},
		{name: "Date format type 2", input: "20060102", output: "20060102", err: nil},
		{name: "Date format type 3", input: "2006-01-02", output: "20060102", err: nil},
		{name: "Date format type 4: empty string", input: "      ", output: "", err: ErrDatatimeMalformed},
		{name: "Date format type 5", input: "2006-01-02T00:00:00Z", output: "20060102", err: nil},
		{name: "Wrong Date format", input: "dummy", output: "", err: ErrDatatimeMalformed},
	}

	//ACT
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := IsoDateFormatter(c.input)
			assert.Equal(t, c.err, err)
		})
	}
}

func TestIsISODate(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		parsed bool
	}{
		{name: "Date format type 1", input: "02.01.2006", parsed: true},
		{name: "Date format type 2", input: "20060102", parsed: true},
		{name: "Date format type 3", input: "2006-01-02", parsed: true},
		{name: "Date format type 4: empty string", input: "      ", parsed: false},
		{name: "Date format type 5", input: "2006-01-02T00:00:00Z", parsed: true},
		{name: "Wrong Date format", input: "dummy", parsed: false},
	}

	for _, tc := range cases {
		assert.Equal(t, tc.parsed, IsISODate(tc.input))
	}
}

func TestISODateStringBeforeToday(t *testing.T) {
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Now().Location())
	cases := []struct {
		name         string
		validateDate string
		expected     bool
		err          error
	}{
		{name: "Date today", validateDate: ToISO8601DateString(today), expected: false, err: nil},
		{name: "Date after 1 day", validateDate: ToISO8601DateString(today.AddDate(0, 0, 1)), expected: false, err: nil},
		{name: "Date after 1 month", validateDate: ToISO8601DateString(today.AddDate(0, 1, 0)), expected: false, err: nil},
		{name: "Date after 1 year", validateDate: ToISO8601DateString(today.AddDate(1, 0, 0)), expected: false, err: nil},
		{name: "Date before 1 day", validateDate: ToISO8601DateString(today.AddDate(0, 0, -1)), expected: true, err: nil},
		{name: "Date before 1 month", validateDate: ToISO8601DateString(today.AddDate(0, -1, 0)), expected: true, err: nil},
		{name: "Date before 1 year", validateDate: ToISO8601DateString(today.AddDate(-1, 0, 0)), expected: true, err: nil},
		{name: "Date malformed", validateDate: "012345678", expected: false, err: ErrDatatimeMalformed},
	}

	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			got, err := ISODateStringBeforeToday(cases[i].validateDate)
			assert.Equal(t, cases[i].err, err)
			assert.Equal(t, cases[i].expected, got)
		})
	}
}

func TestToISO86012004BasicString(t *testing.T) {
	then, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	require.Nil(t, err)

	cases := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{name: "Date #1", input: "02.01.2006", expected: "20060102", err: nil},
		{name: "Date #2", input: "20060102", expected: "20060102", err: nil},
		{name: "Date #3", input: "2006-01-02", expected: "20060102", err: nil},
		{name: "Date #4", input: ToISO86012004BasicString(then), expected: "20060102", err: nil},
		{name: "Date #5", input: "012345678", expected: ToISO86012004BasicString(time.Time{}), err: ErrDatatimeMalformed},
	}

	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			d, err := IsoDateFormatter(cases[i].input)
			assert.Equal(t, cases[i].err, err)
			assert.Equal(t, cases[i].expected, ToISO86012004BasicString(d))
		})
	}
}

func TestToISO8601DateString(t *testing.T) {
	then, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	require.Nil(t, err)

	cases := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{name: "Date #1", input: "02.01.2006", expected: "2006-01-02", err: nil},
		{name: "Date #2", input: "20060102", expected: "2006-01-02", err: nil},
		{name: "Date #3", input: "2006-01-02", expected: "2006-01-02", err: nil},
		{name: "Date #4", input: ToISO8601DateString(then), expected: "2006-01-02", err: nil},
		{name: "Date #5", input: "012345678", expected: ToISO8601DateString(time.Time{}), err: ErrDatatimeMalformed},
	}

	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			d, err := IsoDateFormatter(cases[i].input)
			assert.Equal(t, cases[i].err, err)
			assert.Equal(t, cases[i].expected, ToISO8601DateString(d))
		})
	}
}

func TestToDEFormatDateString(t *testing.T) {
	then, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	require.Nil(t, err)

	cases := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{name: "Date #1", input: "2006-01-02", expected: "02.01.2006", err: nil},
		{name: "Date #2", input: "20060102", expected: "02.01.2006", err: nil},
		{name: "Date #3", input: "2006-01-02", expected: "02.01.2006", err: nil},
		{name: "Date #4", input: ToDEFormatDateString(then), expected: "02.01.2006", err: nil},
		{name: "Date #5", input: "012345678", expected: ToDEFormatDateString(time.Time{}), err: ErrDatatimeMalformed},
	}

	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			d, err := IsoDateFormatter(cases[i].input)
			assert.Equal(t, cases[i].err, err)
			assert.Equal(t, cases[i].expected, ToDEFormatDateString(d))
		})
	}
}

func TestToDEFormatTimeSecondsString(t *testing.T) {
	then, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	require.Nil(t, err)

	cases := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{name: "Date #1", input: "2006-01-02 13:05:02", expected: "02.01.2006 13:05:02", err: nil},
		{name: "Date #2", input: "20060102", expected: "02.01.2006 00:00:02", err: nil},
		{name: "Date #3", input: "2006-01-02 18:25:02", expected: "02.01.2006 18:25:02", err: nil},
		{name: "Date #4", input: ToDEFormatDateString(then), expected: "02.01.2006 00:00:02", err: nil},
		{name: "Date #5", input: "012345678", expected: ToDEFormatDateTimeSecondsString(time.Time{}), err: ErrDatatimeMalformed},
	}

	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			d, err := IsoDateFormatter(cases[i].input)
			assert.Equal(t, cases[i].err, err)
			assert.Equal(t, cases[i].expected, ToDEFormatDateTimeSecondsString(d))
		})
	}
}

func TestToDEFormatDateTimeString(t *testing.T) {
	then, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	require.Nil(t, err)

	cases := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{name: "Date #1", input: "2006-01-02 13:05:02", expected: "02.01.2006 13:05", err: nil},
		{name: "Date #2", input: "20060102", expected: "02.01.2006 00:00", err: nil},
		{name: "Date #3", input: "2006-01-02 18:25:02", expected: "02.01.2006 18:25", err: nil},
		{name: "Date #4", input: ToDEFormatDateString(then), expected: "02.01.2006 00:00", err: nil},
		{name: "Date #5", input: "012345678", expected: ToDEFormatDateTimeString(time.Time{}), err: ErrDatatimeMalformed},
	}

	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			d, err := IsoDateFormatter(cases[i].input)
			assert.Equal(t, cases[i].err, err)
			assert.Equal(t, cases[i].expected, ToDEFormatDateTimeString(d))
		})
	}
}

func TestToISO8601DateTimeString(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{name: "Date #1", input: "02.01.2006", expected: "2006-01-02T00:00:00+0000", err: nil},
		{name: "Date #2", input: "012345678", expected: ToISO8601DateTimeString(time.Time{}), err: ErrDatatimeMalformed},
	}

	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			d, err := IsoDateFormatter(cases[i].input)
			assert.Equal(t, cases[i].err, err)
			assert.Equal(t, cases[i].expected, ToISO8601DateTimeString(d))
		})
	}
}

func Test_GetDateTimeFromToForTonight(t *testing.T) {
	t.Skip("not used, should be refactored")
	now := time.Now()
	cases := []struct {
		name     string
		now      time.Time
		from     string
		to       string
		fromTime time.Time
		toTime   time.Time
		err      error
	}{
		{name: "Interval after EOD overnight",
			now:      time.Date(now.Year(), now.Month(), now.Day(), 19, 0, 0, 0, time.UTC),
			from:     "22:00", // in "Europe/Berlin" TZ
			to:       "03:00", // in "Europe/Berlin" TZ
			fromTime: time.Date(now.Year(), now.Month(), now.Day(), 20, 0, 0, 0, time.UTC),
			toTime:   time.Date(now.Year(), now.Month(), now.Day()+1, 1, 0, 0, 0, time.UTC),
			err:      nil},
		{name: "Interval after EOD daylight",
			now:      time.Date(now.Year(), now.Month(), now.Day(), 19, 0, 0, 0, time.UTC),
			from:     "02:00", // in "Europe/Berlin" TZ
			to:       "22:00", // in "Europe/Berlin" TZ
			fromTime: time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC),
			toTime:   time.Date(now.Year(), now.Month(), now.Day()+1, 20, 0, 0, 0, time.UTC),
			err:      nil},
		{name: "Interval after EOD 24h",
			now:      time.Date(now.Year(), now.Month(), now.Day(), 19, 0, 0, 0, time.UTC),
			from:     "22:00", // in "Europe/Berlin" TZ
			to:       "22:00", // in "Europe/Berlin" TZ
			fromTime: time.Date(now.Year(), now.Month(), now.Day(), 20, 0, 0, 0, time.UTC),
			toTime:   time.Date(now.Year(), now.Month(), now.Day()+1, 20, 0, 0, 0, time.UTC),
			err:      nil},
		{name: "Interval before BOD overnight",
			now:      time.Date(now.Year(), now.Month(), now.Day(), 1, 0, 0, 0, time.UTC),
			from:     "22:00", // in "Europe/Berlin" TZ
			to:       "03:00", // in "Europe/Berlin" TZ
			fromTime: time.Date(now.Year(), now.Month(), now.Day()-1, 20, 0, 0, 0, time.UTC),
			toTime:   time.Date(now.Year(), now.Month(), now.Day(), 1, 0, 0, 0, time.UTC),
			err:      nil},
		{name: "Interval before BOD daylight",
			now:      time.Date(now.Year(), now.Month(), now.Day(), 1, 0, 0, 0, time.UTC),
			from:     "02:00", // in "Europe/Berlin" TZ
			to:       "22:00", // in "Europe/Berlin" TZ
			fromTime: time.Date(now.Year(), now.Month(), now.Day()-1, 0, 0, 0, 0, time.UTC),
			toTime:   time.Date(now.Year(), now.Month(), now.Day(), 20, 0, 0, 0, time.UTC),
			err:      nil},
		{name: "Interval before BOD 24h",
			now:      time.Date(now.Year(), now.Month(), now.Day()+1, 3, 0, 0, 0, time.UTC),
			from:     "02:00", // in "Europe/Berlin" TZ
			to:       "02:00", // in "Europe/Berlin" TZ
			fromTime: time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC),
			toTime:   time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, time.UTC),
			err:      nil},
	}

	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			fromTime, toTime, err := GetDateTimeFromToForTonight(cases[i].now, cases[i].from, cases[i].to)
			assert.Equal(t, cases[i].err, err)
			assert.Equal(t, cases[i].fromTime, fromTime.UTC())
			assert.Equal(t, cases[i].toTime, toTime.UTC())
		})
	}
}

func Test_GetDateTimeFromTo(t *testing.T) {
	now := time.Now()
	cases := []struct {
		name     string
		now      time.Time
		from     string
		to       string
		fromTime time.Time
		toTime   time.Time
		err      error
	}{
		{name: "Interval after EOD overnight, negative",
			now:      time.Date(now.Year(), now.Month(), now.Day(), 23, 0, 0, 0, time.UTC),
			from:     "11:00",
			to:       "03:00",
			fromTime: time.Time{},
			toTime:   time.Time{},
			err:      ErrWrongTimeInterval},
		{name: "Interval after EOD daylight",
			now:      time.Date(now.Year(), now.Month(), now.Day(), 19, 0, 0, 0, time.UTC),
			from:     "02:00",
			to:       "22:00",
			fromTime: time.Date(now.Year(), now.Month(), now.Day(), 2, 0, 0, 0, time.UTC),
			toTime:   time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 0, 0, time.UTC),
			err:      nil},
		{name: "Interval after EOD overnight",
			now:      time.Date(now.Year(), now.Month(), now.Day(), 23, 0, 0, 0, time.UTC),
			from:     "02:00",
			to:       "22:00",
			fromTime: time.Date(now.Year(), now.Month(), now.Day(), 2, 0, 0, 0, time.UTC),
			toTime:   time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 0, 0, time.UTC),
			err:      nil},
		{name: "Interval after EOD 0h overnight",
			now:      time.Date(now.Year(), now.Month(), now.Day(), 23, 0, 0, 0, time.UTC),
			from:     "22:00",
			to:       "22:00",
			fromTime: time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 0, 0, time.UTC),
			toTime:   time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 0, 0, time.UTC),
			err:      nil},
		{name: "Interval before BOD overnight, negative",
			now:      time.Date(now.Year(), now.Month(), now.Day(), 1, 0, 0, 0, time.UTC),
			from:     "22:00",
			to:       "03:00",
			fromTime: time.Time{},
			toTime:   time.Time{},
			err:      ErrWrongTimeInterval},
		{name: "Interval before BOD overnight",
			now:      time.Date(now.Year(), now.Month(), now.Day(), 1, 0, 0, 0, time.UTC),
			from:     "02:00",
			to:       "22:00",
			fromTime: time.Date(now.Year(), now.Month(), now.Day(), 2, 0, 0, 0, time.UTC),
			toTime:   time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 0, 0, time.UTC),
			err:      nil},
		{name: "Interval before BOD daylight",
			now:      time.Date(now.Year(), now.Month(), now.Day(), 8, 0, 0, 0, time.UTC),
			from:     "02:00",
			to:       "22:00",
			fromTime: time.Date(now.Year(), now.Month(), now.Day(), 2, 0, 0, 0, time.UTC),
			toTime:   time.Date(now.Year(), now.Month(), now.Day(), 22, 0, 0, 0, time.UTC),
			err:      nil},
		{name: "Interval before BOD 0h",
			now:      time.Date(now.Year(), now.Month(), now.Day(), 3, 0, 0, 0, time.UTC),
			from:     "02:00",
			to:       "02:00",
			fromTime: time.Date(now.Year(), now.Month(), now.Day(), 2, 0, 0, 0, time.UTC),
			toTime:   time.Date(now.Year(), now.Month(), now.Day(), 2, 0, 0, 0, time.UTC),
			err:      nil},
	}

	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			fromTime, toTime, err := GetDateTimeFromTo(cases[i].now, cases[i].from, cases[i].to)
			assert.Equal(t, cases[i].err, err)
			assert.Equal(t, cases[i].fromTime, fromTime.UTC())
			assert.Equal(t, cases[i].toTime, toTime.UTC())
		})
	}
}
