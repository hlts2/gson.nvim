package command

import "strings"

// GsonViewerCommand is gson-viewer.nvim command
func GsonViewerCommand(args []string) (string, error) {
	return "Hello GsonViewer, args: " + strings.Join(args, " "), nil
}
