#!/usr/bin/env bash

gcc -c -fPIC f.c
gcc -c -fPIC g.c
gcc -shared -Wl,-soname,libfg.so.1 f.o g.o -o libfg.so.1
ln -f -s libfg.so.1 libfg.so
gcc -c main.c
gcc -L. main.o -lfg -lc -o prog-shared

# LD_LIBRARY_PATH=. ./prog-shared

