// Package helps parse http headers.
package httphead

import (
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

// ContentType describes "Content-Type"
var ContentType = struct {
	ApplicationZip         string
	TextCsv                string
	TextPlain              string
	ApplicationOctetStream string
	ApplicationJSON        string
	MultipartFormData      string
}{
	"application/zip",
	"text/csv",
	"text/plain",
	"application/octet-stream",
	"application/json",
	"multipart/form-data",
}

// GetContentTypeByFileName returns "Content-Type" by file name
func GetContentTypeByFileName(name string) string {
	switch strings.ToLower(filepath.Ext(name)) {
	case ".zip":
		return ContentType.ApplicationZip
	case ".csv":
		return ContentType.TextCsv
	default:
		return ContentType.TextPlain // if no extension found
	}
}

// HasContentType determines whether the request `content-type` includes a server-acceptable mime-type
func HasContentType(r *http.Request, mimetype string) bool {
	contentType := r.Header.Get("Content-type")
	if contentType == "" {
		return mimetype == ContentType.ApplicationOctetStream
	}
	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype {
			return true
		}
	}
	return false
}
//
//const (
//	ApplicationZip  ContentType = "application/zip"
//	ApplicationOctetStream ContentType = "application/octet-stream"
//	ApplicationJSON ContentType = "application/json"
//	MultipartFormData ContentType = "multipart/form-data"
//	TextCsv ContentType = "text/csv"
//	TextPlain ContentType = "text/plain"
//)
//
////with this I cannot make strVal == ApplicationZip by default because ApplicationZip is not string
////look at load-balancer to use this pattern
//type ContentType string
//
////type Header struct {
////	ContentType	ContentType
////	Body     string
////}
//
