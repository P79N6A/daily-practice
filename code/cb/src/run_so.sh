#!/usr/bin/env bash

if [ ! -f "prog-shared" ]; then 
	./compile_so.sh
fi

LD_LIBRARY_PATH=. ./prog-shared
