package gsonnvim

import "github.com/neovim/go-client/nvim"

const (

	// StartHighlightMark is named of highlight start mark
	StartHighlightMark = "<"

	// EndHighlightMark is named of highlight end mark
	EndHighlightMark = ">"
)

// HighlightMark is highlight position struct
type HighlightMark struct {
	start [2]int
	end   [2]int
}

// NewHighlightMark returns HighlightMark instance
func NewHighlightMark(v *nvim.Nvim, buf nvim.Buffer) (*HighlightMark, error) {
	start, err := v.BufferMark(buf, StartHighlightMark)
	if err != nil {
		return nil, err
	}

	end, err := v.BufferMark(buf, EndHighlightMark)
	if err != nil {
		return nil, err
	}

	hm := &HighlightMark{}

	hm.start = start
	hm.end = end

	return hm, nil
}

// AdjustMarkWithRange adjusts mark with range
func (hm *HighlightMark) AdjustMarkWithRange(startLineNum, endLineNum int) {
	if hm.start[0] != startLineNum || hm.end[0] != endLineNum {
		hm.start[0] = startLineNum
		hm.start[1] = 0

		hm.end[0] = endLineNum
		hm.end[1] = 0
	}
}

// GetStart returns line number and line position of highlight start mark.
func (hm *HighlightMark) GetStart() [2]int {
	return hm.start
}

// GetStartLineNum returns line number of highlight start mark
func (hm *HighlightMark) GetStartLineNum() int {
	return hm.start[0]
}

// GetStartPos returns line position of highlight start mark
func (hm *HighlightMark) GetStartPos() int {
	return hm.start[1]
}

// GetEnd returns line number and line position of highlight end mark.
func (hm *HighlightMark) GetEnd() [2]int {
	return hm.end
}

// GetEndLineNum returns line number of highlight end mark
func (hm *HighlightMark) GetEndLineNum() int {
	return hm.end[0]
}

// GetEndPos returns line position of highlight end mark
func (hm *HighlightMark) GetEndPos() int {
	return hm.end[1]
}
