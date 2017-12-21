From ubuntu:16.04
MAINTAINER Kai Qiang Wu

WORKDIR /tempwork

ADD vimrc /tempwork

RUN sed -i 's/archive.ubuntu.com/mirrors.aliyun.com/g'  /etc/apt/sources.list && \
  apt-get update && apt-get install -y gcc git python python-dev ruby-dev ncurses-dev \
  build-essential make cmake ack-grep exuberant-ctags wget xz-utils libboost1.58-all-dev

RUN git clone https://github.com/vim/vim.git && \
  cd vim/ && ./configure --with-features=huge --enable-pythoninterp --enable-rubyinterp --enable-luainterp --enable-perlinterp --with-python-config-dir=/usr/lib/python2.7/config-x86_64-linux-gnu/  --enable-cscope --prefix=/usr && \
  make && make install

RUN cp vimrc ~/.vimrc && mkdir -p ~/.vim/bundle && \
  git clone https://github.com/gmarik/Vundle.vim.git ~/.vim/bundle/Vundle.vim && \
  git clone https://github.com/altercation/vim-colors-solarized.git  ~/.vim/bundle/vim-colors-solarized && \
  git clone --depth 1 https://github.com/Valloric/YouCompleteMe.git ~/.vim/bundle/YouCompleteMe && \
  vim +PluginInstall +qall && \
  cd ~/.vim/bundle/YouCompleteMe && git submodule update --init --recursive && \
  ./install.py

RUN wget http://releases.llvm.org/5.0.0/clang+llvm-5.0.0-linux-x86_64-ubuntu16.04.tar.xz && \
  xz -d clang+llvm-5.0.0-linux-x86_64-ubuntu16.04.tar.xz && \
  mkdir clang+llvm && tar -xvf clang+llvm-5.0.0-linux-x86_64-ubuntu16.04.tar -C clang+llvm && \
  mkdir ycm_build && cd ycm_build && \
  cmake -G "Unix Makefiles" -DUSE_SYSTEM_BOOST=ON -DPATH_TO_LLVM_ROOT=../clang+llvm/clang+llvm-5.0.0-linux-x86_64-ubuntu16.04/ .  ~/.vim/bundle/YouCompleteMe/third_party/ycmd/cpp && \
  cd ~/.vim/bundle/YouCompleteMe && ./install.py --clang-completer && rm -rf /tempwork
