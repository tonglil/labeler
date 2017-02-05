package versioning

import (
	"fmt"
	"io"
)

// Deliberately uninitialized so it can be set during the build process.
var version string

// Set sets the version.
func Set(v string) {
	version = v
}

// Get returns the version.
func Get() string {
	if version != "" {
		return version
	}

	return "unknown"
}

// Write the version to a writer.
func Write(w io.Writer) {
	fmt.Fprintf(w, "version %s\n", Get())
}
