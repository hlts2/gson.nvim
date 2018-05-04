package util

import (
	"bytes"
	"unsafe"
)

// SliceByCharIndex return range value based on index of character
func SliceByCharIndex(s []byte, sIdx, eIdx int) []byte {
	if !checkRange(s, sIdx, eIdx) {
		return s
	}

	str := string(bytes.Runes(s)[sIdx:eIdx])
	return *(*[]byte)(unsafe.Pointer(&str))
}

func checkRange(s []byte, sIdx, eIdx int) bool {
	if !(eIdx > sIdx && sIdx > -1) {
		return false
	}

	if len(s) < eIdx {
		return false
	}

	return true
}
