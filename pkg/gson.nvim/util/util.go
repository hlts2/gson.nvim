package gsonnvim

import (
	"path/filepath"
	"unicode/utf8"
)

// IsJSONFile returns true if the path is json file.
func IsJSONFile(path string) bool {
	fileN := filepath.Base(path)

	charCnt := utf8.RuneCountInString(fileN)
	if charCnt < 5 {
		return false
	}

	return fileN[len(fileN)-5:len(fileN)] == ".json"
}
