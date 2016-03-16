
Good Golang+VIM IDE Tool to Help Daily Life
==========================================================

As a developer, I need quickly setup golang IDE environment.

> **Note**: The Dockerfile part of it is based on
> https://hub.docker.com/r/mbrt/golang-vim-dev/~/dockerfile/

## How to Run it?

Build the image:

	docker build -t golangdevenv:1.0  .

Start use it:
	
	docker run -it --rm -v `pwd`:/go bash

## TODO List

Track any new vim plugin if can improve daily work.

## Bugs

Latest vim not work with vimproc plugin.

Not know now :)


