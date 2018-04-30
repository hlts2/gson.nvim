package command

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/hlts2/gson"
	"github.com/neovim/go-client/nvim"
)

// Fmt is json format command
func (gn *GsonNvim) Fmt(v *nvim.Nvim, args []string) error {
	vBuf, err := v.CurrentBuffer()
	if err != nil {
		return err
	}

	str, _ := v.BufferName(vBuf)
	return errors.New(fmt.Sprint("%v", str))

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
