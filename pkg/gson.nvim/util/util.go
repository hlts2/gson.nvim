package util

import (
	"path/filepath"
	"unicode/utf8"
)

// IsJSONFile returns true if the path is json file.
func IsJSONFile(path string) bool {
	fileN := filepath.Base(path)

	charCnt := utf8.RuneCountInString(fileN)

	// The smallest unit constituting the json file
	// For example, if 「a.json」 is the smallest unit
	if charCnt < 6 {
		return false
	}

	return fileN[len(fileN)-5:len(fileN)] == ".json"
}
