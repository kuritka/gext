package httphead

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetContentTypeByFileName(t *testing.T) {
	cases := []struct {
		name     string
		fileName string
		expected string
	}{
		{name: "Test Content-type CSV #1", fileName: "example.csv", expected: ContentType.TextCsv},
		{name: "Test Content-type CSV #2", fileName: "example.CSV", expected: ContentType.TextCsv},
		{name: "Test Content-type ZIP #1", fileName: "example.zip", expected: ContentType.ApplicationZip},
		{name: "Test Content-type ZIP #2", fileName: "example.ZIP", expected: ContentType.ApplicationZip},
		{name: "Test Content-type TXT #1", fileName: "example.txt", expected: ContentType.TextPlain},
		{name: "Test Content-type TXT #2", fileName: "example.TXT", expected: ContentType.TextPlain},
	}

	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			got := GetContentTypeByFileName(cases[i].fileName)
			assert.Equal(t, cases[i].expected, got)
		})
	}
}

func TestHasContentType(t *testing.T) {
	cases := []struct {
		name         string
		setType      string
		validateType string
		expected     bool
	}{
		{name: "Test Content-type is 'text/csv'", setType: ContentType.TextCsv, validateType: ContentType.TextCsv, expected: true},
		{name: "Test Content-type is not 'text/csv'", setType: ContentType.TextCsv, validateType: ContentType.TextPlain, expected: false},
		{name: "Test Content-type is 'application/octet-stream'", setType: ContentType.ApplicationOctetStream, validateType: ContentType.ApplicationOctetStream, expected: true},
		{name: "Test Content-type is not 'application/octet-stream'", setType: ContentType.TextCsv, validateType: ContentType.ApplicationOctetStream, expected: false},
		{name: "Test Content-type is default 'application/octet-stream'", setType: "", validateType: ContentType.ApplicationOctetStream, expected: true},
		{name: "Test Content-type is 'multipart/form-data'", setType: ContentType.MultipartFormData, validateType: ContentType.MultipartFormData, expected: true},
		{name: "Test Content-type is not 'multipart/form-data'", setType: ContentType.TextCsv, validateType: ContentType.MultipartFormData, expected: false},
	}

	r, err := http.NewRequest(http.MethodGet, "/notifications", nil)
	require.Nil(t, err)

	for i := range cases {
		t.Run(cases[i].name, func(t *testing.T) {
			r.Header.Set(http.CanonicalHeaderKey("Content-type"), cases[i].setType)
			got := HasContentType(r, cases[i].validateType)
			assert.Equal(t, cases[i].expected, got)
		})
	}
}
