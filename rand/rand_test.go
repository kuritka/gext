package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomUUID(t *testing.T) {
	cases := []string{
		"^[0-9a-z]{8}-([0-9a-z]{4}-){3}[0-9a-z]{12}$", // 7a13663d-2f95-42ba-acd6-2c6fed94fca5
	}
	for _, tc := range cases {
		uuid, err := GenerateRandomUUID()
		require.Nil(t, err)
		assert.Regexp(t, tc, uuid)
	}
}

func TestGenerateRandomString(t *testing.T) {
	cases := []int{0, -1, 15}

	for _, size := range cases {
		randStr := GenerateRandomString(size)
		if size < 1 {
			assert.Equal(t, 0, len(randStr))
		} else {
			assert.Equal(t, size, len(randStr))
		}
	}
}

func TestParseUUID(t *testing.T) {
	cases := []struct {
		name       string
		sampleUUID string
		expected   string
		err        error
	}{
		{name: "ParseUUID", sampleUUID: "85B96F37-03A6-4731-A8A3-E1629384B3F5", expected: "85b96f37-03a6-4731-a8a3-e1629384b3f5", err: nil},
		{name: "ParseUUID", sampleUUID: "85B96F3703A64731A8A3E1629384B3F5", expected: "85b96f37-03a6-4731-a8a3-e1629384b3f5", err: nil},
		{name: "ParseUUID", sampleUUID: "00000000-0000-0000-0000-000000000000", expected: "00000000-0000-0000-0000-000000000000", err: nil},
		{name: "ParseUUID", sampleUUID: "497D106E0-039C-4FC7-BD59-955549E048A1", expected: "", err: errors.New("invalid UUID length: 37")},
		{name: "ParseUUID", sampleUUID: "497D106E00-039C-4FC7-BD59-955549E048A1", expected: "", err: errors.New("invalid UUID format")},
		{name: "ParseUUID", sampleUUID: "", expected: "", err: errors.New("invalid UUID length: 0")},
	}
	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			got, err := ParseUUID(cases[i].sampleUUID)
			require.Equal(t, cases[i].err, err)
			assert.Equal(t, cases[i].expected, got)
		})
	}
}

func TestGenerateRandomNumber(t *testing.T) {
	cases := []struct {
		name string
		min  int
		max  int
	}{
		{name: "GenerateRandomNumber", min: 0, max: 1},
		{name: "GenerateRandomNumber", min: 0, max: 9},
		{name: "GenerateRandomNumber", min: 999999, max: 9999999},
	}
	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			got := GenerateRandomNumber(cases[i].min, cases[i].max)
			assert.True(t, got >= cases[i].min && got <= cases[i].max)
		})
	}
}
