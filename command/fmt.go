package command

import (
	"bytes"

	"github.com/hlts2/gson"
	gsonnvim "github.com/hlts2/gson.nvim/pkg/gson_nvim"
	"github.com/neovim/go-client/nvim"
)

const (

	// JSONIndent is indent of JSON
	JSONIndent = "  "

	// JSONIndentPrefix is prefix of JSON indent
	JSONIndentPrefix = ""
)

// CR is carriage return
var CR = []byte{10}

// Fmt is json format command
func (gn *GsonNvim) Fmt(v *nvim.Nvim, args []string, rng [2]int) error {
	vBuf, err := v.CurrentBuffer()
	if err != nil {
		return err
	}

	highlightMark, err := gsonnvim.NewHighlightMark(v, vBuf)
	if err != nil {
		return err
	}

	highlightMark.AdjustMarkWithRange(rng[0], rng[1])

	rangeArea, err := gsonnvim.NewRangeArea(v, vBuf, highlightMark.GetStart(), highlightMark.GetEnd())
	if err != nil {
		return err
	}

	g, err := gson.NewGsonFromByte(bytes.Join(rangeArea.GetSelectedLines(), CR))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := g.Indent(&buf, JSONIndentPrefix, JSONIndent); err != nil {
		return err
	}

	startLineNum := highlightMark.GetStartLineNum()
	endLineNum := highlightMark.GetEndLineNum()

	v.SetBufferLines(vBuf, startLineNum-1, endLineNum, true, bytes.Split(append(append(rangeArea.GetUnselectedStartLine(), buf.Bytes()...), rangeArea.GetUnselectedEndLine()...), CR))

	return nil
}
