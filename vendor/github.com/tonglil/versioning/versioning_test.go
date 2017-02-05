package versioning

import (
	"testing"
)

// TestSetGet tests both Set and Get functions.
func TestSetGet(t *testing.T) {
	cases := []struct {
		version string
		expect  string
	}{
		{"", "unknown"},
		{"1.2.3", "1.2.3"},
	}

	for _, tc := range cases {
		Set(tc.version)
		actual := Get()

		if actual != tc.expect {
			t.Fatalf("expected %s for %s, got %s", tc.expect, tc.version, actual)
		}
	}
}
