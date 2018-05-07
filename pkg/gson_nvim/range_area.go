package gsonnvim

import (
	"unsafe"

	"github.com/neovim/go-client/nvim"
)

// RangeArea is range lines structor
type RangeArea struct {
	selectedLines       [][]byte
	unselectedStartLine []byte
	unselectedEndLine   []byte
}

// NewRangeArea returns RangeArea instance
func NewRangeArea(v *nvim.Nvim, buf nvim.Buffer, sMark [2]int, eMark [2]int) (*RangeArea, error) {
	lines, err := v.BufferLines(buf, sMark[0]-1, eMark[0], true)
	if err != nil {
		return nil, err
	}

	ra := &RangeArea{}

	ra.unselectedStartLine = getUnselectedStartLine(lines, sMark[1])
	ra.unselectedEndLine = getUnselectedEndLine(lines, eMark[1])
	ra.selectedLines = getSelectedLines(lines, sMark, eMark)

	return ra, nil
}

// GetUnselectedStartLine returns unselected start line
func (ra *RangeArea) GetUnselectedStartLine() []byte {
	return ra.unselectedStartLine
}

func getUnselectedStartLine(lines [][]byte, pos int) []byte {
	if pos == 0 {
		return []byte{}
	}

	return lines[0][:pos]
}

// GetUnselectedEndLine returns unselected end line
func (ra *RangeArea) GetUnselectedEndLine() []byte {
	return ra.unselectedEndLine
}

func getUnselectedEndLine(lines [][]byte, pos int) []byte {
	if pos == 0 {
		return []byte{}
	}

	endLen := len(lines) - 1

	return lines[endLen][pos:]
}

func (ra *RangeArea) GetSelectedLines() [][]byte {
	return ra.selectedLines
}

func getSelectedLines(lines [][]byte, sMark [2]int, eMark [2]int) [][]byte {
	sLineNum, eLineNum := sMark[0], eMark[0]
	sPos, ePos := sMark[1], eMark[1]

	endLen := len(lines) - 1

	if ePos == 0 {
		ePos = len(lines[endLen])
	}

	if sLineNum == eLineNum {
		lines[0] = lines[0][sPos:ePos]
	} else {
		lines[0] = lines[0][sPos:]
		lines[endLen] = lines[endLen][:ePos]
	}

	return lines
}

func getNextPos(line []byte, pos int) int {
	if pos == 0 {
		return len(line)
	}

	exist := false
	for byteIdx := range *(*string)(unsafe.Pointer(&line)) {
		if pos == byteIdx {
			exist = true
			continue
		}

		if exist {
			return byteIdx
		}
	}

	return len(line)
}
