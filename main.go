package main

import (
	"github.com/hlts2/gson-viewer.nvim/command"
	"github.com/neovim/go-client/nvim/plugin"
)

func main() {
	plugin.Main(func(p *plugin.Plugin) error {
		p.HandleFunction(&plugin.FunctionOptions{Name: "GsonViewer"}, command.GsonViewerCommand)
		return nil
	})
}
