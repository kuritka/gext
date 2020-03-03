// Package utils implements helpers.
package guard

import (
	"fmt"
	"os"
)

// Must exit on error.
func Must(err error) {
	if err == nil {
		return
	}

	fmt.Printf("ERROR: %+v\n", err)
	os.Exit(1)
}
