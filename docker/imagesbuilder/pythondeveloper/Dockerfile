From ubuntu:16.04
MAINTAINER Kai Qiang Wu

ADD vimrc /

RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g'  /etc/apt/sources.list && \
  apt-get update && apt-get install -y sudo vim git python python2.7-dev build-essential cmake ack-grep exuberant-ctags && \
  adduser dev --disabled-password --gecos "" && \
  echo "ALL            ALL = (ALL) NOPASSWD: ALL" >> /etc/sudoers && \
  chown -R dev:dev /home/dev

USER dev
ENV HOME /home/dev

RUN mkdir -p ~/.vim/bundle && \
  git clone https://github.com/gmarik/Vundle.vim.git ~/.vim/bundle/Vundle.vim && \
  cd ~/.vim/bundle && \
  git clone --depth 1 https://github.com/Valloric/YouCompleteMe.git && \
  git clone --depth 1 https://github.com/airblade/vim-gitgutter.git && \
  git clone --depth 1 https://github.com/kien/ctrlp.vim.git && \
  git clone --depth 1 https://github.com/scrooloose/nerdtree.git && \
  git clone --depth 1 https://github.com/tpope/vim-fugitive.git && \
  git clone --depth 1 https://github.com/vim-airline/vim-airline-themes.git && \
  git clone --depth 1 https://github.com/vim-airline/vim-airline.git && \
  git clone --depth 1 https://github.com/scrooloose/syntastic.git && \
  git clone --depth 1 https://github.com/easymotion/vim-easymotion.git && \
  git clone --depth 1 https://github.com/mileszs/ack.vim.git && \
  git clone --depth 1 https://github.com/ervandew/supertab.git && \
  git clone --depth 1 https://github.com/majutsushi/tagbar.git && \
  git clone --depth 1 https://github.com/mbbill/undotree.git && \
  git clone --depth 1 https://github.com/godlygeek/tabular.git && \
  git clone --depth 1 https://github.com/plasticboy/vim-markdown.git && \
  git clone --depth 1 https://github.com/elzr/vim-json.git && \
  git clone --depth 1 https://github.com/ntpeters/vim-better-whitespace.git

RUN cp /vimrc ~/.vimrc && \
  vim +PluginInstall +qall && \
  cd ~/.vim/bundle/YouCompleteMe && git submodule update --init --recursive && \
  ./install.py --clang-completer

RUN rm -rf Vundle.vim/.git tagbar/.git nerdtree/.git vim-easymotion/.git \
    vim-fugitive/.git vim-airline/.git vim-airline-themes/.git syntastic/.git \
    supertab/.git ack.vim/.git ctrlp.vim/.git YouCompleteMe/.git vim-gitgutter/.git \
    undotree/.git tabular/.git vim-markdown/.git vim-json/.git vim-better-whitespace/.git
