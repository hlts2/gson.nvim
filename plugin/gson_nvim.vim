if exists('g:loaded_gson')
  finish
endif
let g:loaded_gson = 1

function! s:RequireGson(host) abort
  return jobstart(['gson.nvim'], {'rpc': v:true})
endfunction

call remote#host#Register('gson.nvim', '0', function('s:RequireGson'))
call remote#host#RegisterPlugin('gson.nvim', '0', [
    \ {'type': 'command', 'name': 'GsonFmt', 'sync': 1, 'opts': {'nargs': '?', 'range': '%'}},
    \ ])

