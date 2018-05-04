package util

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	tests := []struct {
		s        []byte
		sIdx     int
		eIdx     int
		expected []byte
	}{
		{
			s:        []byte("hello"),
			sIdx:     0,
			eIdx:     5,
			expected: []byte("hello"),
		},
		{
			s:        []byte("hello"),
			sIdx:     1,
			eIdx:     5,
			expected: []byte("ello"),
		},
		{
			s:        []byte("hello ヒロト"),
			sIdx:     1,
			eIdx:     7,
			expected: []byte("ello ヒ"),
		},
		{
			s:        []byte("hello"),
			sIdx:     -1,
			eIdx:     3,
			expected: []byte("hello"),
		},
		{
			s:        []byte("hello"),
			sIdx:     1,
			eIdx:     -3,
			expected: []byte("hello"),
		},
		{
			s:        []byte("hello"),
			sIdx:     -1,
			eIdx:     -3,
			expected: []byte("hello"),
		},
		{
			s:        []byte("hello"),
			sIdx:     10,
			eIdx:     11,
			expected: []byte("hello"),
		},
		{
			s:        []byte("hello"),
			sIdx:     3,
			eIdx:     2,
			expected: []byte("hello"),
		},
		{
			s:        []byte("hello"),
			sIdx:     1,
			eIdx:     100,
			expected: []byte("hello"),
		},
	}

	for i, test := range tests {
		got := SliceByCharIndex(test.s, test.sIdx, test.eIdx)

		if !reflect.DeepEqual(test.expected, got) {
			t.Errorf("i = %d Slice(s, sIdx, eIdx) expected: %v, got: %v", i, test.expected, got)
		}
	}
}
