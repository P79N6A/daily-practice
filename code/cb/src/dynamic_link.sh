#!/usr/bin/env bash

gcc -c main.c
gcc -c f.c
gcc -v main.o f.o -lc -o prog_dynamic.exe

