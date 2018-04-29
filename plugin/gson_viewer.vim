if exists('g:loaded_gson-viewer')
  finish
endif
let g:loaded_gson-viewer = 1

function! s:RequireGsonViewer(host) abort
  return jobstart(['gson-viewer.nvim'], {'rpc': v:true})
endfunction

call remote#host#Register('gson-viewer.nvim', '0', function('s:RequireGsonViewer'))
call remote#host#RegisterPlugin('gson-viewer.nvim', '0', [
    \ {'type': 'function', 'name': 'GsonViewer', 'sync': 1, 'opts': {}},
    \ ])

