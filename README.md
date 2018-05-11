# gson.nvim  [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![Go Report Card](https://goreportcard.com/badge/github.com/hlts2/gson.nvim)](https://goreportcard.com/report/github.com/hlts2/gson.nvim) [![Join the chat at https://gitter.im/hlts2/gson.nvim](https://badges.gitter.im/hlts2/gson.nvim.svg)](https://gitter.im/hlts2/gson.nvim?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

gson.nvim is a Vim Plugin for JSON.

## Require

- NeoVim( >= 0.2.0)
- go (>= 1.8)

## Install

### 1. Install Neovim
See Neovim wiki.

- [Installing Neovim](https://github.com/neovim/neovim/wiki/Installing-Neovim)
- [Following HEAD](https://github.com/neovim/neovim/wiki/Following-HEAD)
- [Building](https://github.com/neovim/neovim/wiki/Building-Neovim)

### 2. Install go and set up
See go wiki

- [Install and Setup](https://github.com/golang/go/wiki#working-with-go)

### 3. Install Plugin
```vim
" dein.vim
call dein#add('hlts2/gson.nvim', {'build': 'make'})

" NeoBundle
NeoBundle 'hlts2/gson.nvim', {'build': {'unix': 'make'}}

" vim-plug
Plug 'hlts2/gson.nvim', { 'do': 'make'}
```

## Usage
### GosonFmt
- Format JSON

```vim
:GsonFmt
```

#### Demo
<img src="./demo/fmt_demo_1.gif" width="750">

- Format JSON of selection

### Demo
<img src="./demo/fmt_demo_2.gif" width="750">
