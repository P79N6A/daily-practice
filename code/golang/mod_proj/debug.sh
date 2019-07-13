#!/usr/bin/env bash

if [ ! -f ~/.air ]; then
    curl -fLo ~/.air \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/darwin/air
    chmod +x ~/.air
fi

~/.air