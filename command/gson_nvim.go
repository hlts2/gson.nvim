package command

import "errors"

// ErrNotJSONFile represents not JSON file error
var ErrNotJSONFile = errors.New("Not JSON file")

// GsonNvim is GsonNvim base struct
type GsonNvim struct{}
