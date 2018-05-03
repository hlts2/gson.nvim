package command

import (
	"bytes"

	"github.com/hlts2/gson"
	"github.com/hlts2/gson.nvim/pkg/gson.nvim/util"
	"github.com/neovim/go-client/nvim"
)

// Fmt is json format command
func (gn *GsonNvim) Fmt(v *nvim.Nvim, args []string, rng [2]int) error {
	vBuf, err := v.CurrentBuffer()
	if err != nil {
		return err
	}

	path, _ := v.BufferName(vBuf)
	if !util.IsJSONFile(path) {
		return ErrNotJSONFile
	}

	reader := nvim.NewBufferReader(v, vBuf)

	g, err := gson.NewGsonFromReader(reader)
	if err != nil {
		return err
	}

	vBufCount, err := v.BufferLineCount(vBuf)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	err = g.Indent(&buf, "", "  ")
	if err != nil {
		return err
	}

	v.SetBufferLines(vBuf, 0, vBufCount, true, bytes.Split(buf.Bytes(), []byte("\n")))

	return nil
}
