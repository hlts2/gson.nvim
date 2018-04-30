package util

import (
	"testing"
)

func TestIsJSONFile(t *testing.T) {
	tests := []struct {
		path     string
		expected bool
	}{
		{
			path:     "/hoge/aaa/test.json",
			expected: true,
		},
		{
			path:     "test.json",
			expected: true,
		},
		{
			path:     "/hoge/aaa/test.jso",
			expected: false,
		},
		{
			path:     "/hoge/aaa/test.go",
			expected: false,
		},
		{
			path:     "/test.json/fdfdf",
			expected: false,
		},
		{
			path:     "",
			expected: false,
		},
	}

	for _, test := range tests {
		got := IsJSONFile(test.path)

		if test.expected != got {
			t.Errorf("IsJSONFile(path) expected: %v, got: %v", test.expected, got)
		}
	}
}
