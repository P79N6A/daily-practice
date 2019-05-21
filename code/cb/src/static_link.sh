#!/usr/bin/env bash

gcc -c main.c
gcc -c f.c
gcc -v -static main.o f.o -o prog_static.exe
