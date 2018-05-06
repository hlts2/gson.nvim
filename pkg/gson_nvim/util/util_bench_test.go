package util

import (
	"testing"
)

func BenchmarkSliceByCharIndex(b *testing.B) {
	s := []byte("gson.nvim is a Vim Plugin for JSON")
	for i := 0; i < b.N; i++ {
		_ = SliceByCharIndex(s, 10, 24)
	}
}
