set nocompatible              " be iMproved, required
filetype off                  " required

" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
" alternatively, pass a path where Vundle should install plugins
" call vundle#begin('~/some/path/here')
"
" let Vundle manage Vundle, required
Plugin 'VundleVim/Vundle.vim'

" Keep Plugin commands between vundle#begin/end.
Plugin 'Shougo/unite.vim'

Plugin 'fatih/vim-go'

Plugin 'majutsushi/tagbar'

Plugin 'scrooloose/nerdtree'

Plugin 'Shougo/neocomplete.vim'

" All of your Plugins must be added before the following line
call vundle#end()            " required
filetype plugin indent on    " required

execute pathogen#infect()
syntax on
filetype plugin indent on
nmap <F8> :TagbarToggle<CR>
map <C-n> :NERDTreeToggle<CR>
let g:neocomplete#enable_at_startup = 1

nnoremap <silent> ,g :<C-u>Unite grep:. -buffer-name=search-buffer<CR>
if executable('pt')
 let g:unite_source_grep_command = 'pt'
 let g:unite_source_grep_default_opts = '--nogroup --nocolor'
 let g:unite_source_grep_recursive_opt = ''
 let g:unite_source_grep_encoding = 'utf-8'
endif
