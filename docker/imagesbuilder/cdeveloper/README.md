# C/C++ VIM IDE

This IDE is based on https://github.com/yangyangwithgnu/use_vim_as_ide/,
it aslo adds some personal magic code to make it work with docker.
With this docker way, it can be used to set up personal C/C++ IDE quickly.


**Note** `Make tool` or `Code static analysis` is not integrated in this IDE.

## How to make it work?

Build with following:

```
    docker build -t mycdev:1.0  .
```

And start use it with following:

```
    docker run -it -v `pwd`:/home/dev mycdev:1.0 bash
```

## TODO

Much more for you to explore
