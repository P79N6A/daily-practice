#!/bin/bash

# remove all <none> images
docker rmi $(docker images -f "dangling=true" -q)

# remove all exited container
docker rm $(docker ps -q -f status=exited)