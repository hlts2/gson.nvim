package main

import (
	"github.com/hlts2/gson.nvim/command"
	"github.com/neovim/go-client/nvim/plugin"
)

func main() {
	gn := &command.GsonNvim{}

	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleCommand(&plugin.CommandOptions{Name: "GsonFmt"}, gn.Fmt)
		return nil
	})
}
