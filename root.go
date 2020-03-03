// Package gext is a set of packages that provide many utilities and helpers.
//
// gext contains following packages:
//
// The data package provides extended data operations over maps and slices.
//
// The date package formats and parse date-time values.
//
// The env provides easy access to environment variables.
//
// The guard throws errors or panics when error occur
//
// The httphead helps parse http headers
//
// The log wraps zerolog logger and provides standard log functionality
//
// The rand helps with generating random numbers and guids
//
// The reflection provides reflection helpers over structures

package gext

// blank imports help docs.
import (
	// data package
	_ "github.com/kuritka/gext/data"
	// date package
	_ "github.com/kuritka/gext/date"
	// env package
	_ "github.com/kuritka/gext/env"
	// guard package
	_ "github.com/kuritka/gext/guard"
	// httphead package
	_ "github.com/kuritka/gext/httphead"
	// log package
	_ "github.com/kuritka/gext/log"
	// parser package
	_ "github.com/kuritka/gext/parser"
	// rand package
	_ "github.com/kuritka/gext/rand"
	// rand package
	_ "github.com/kuritka/gext/reflection"
)