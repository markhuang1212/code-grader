#!/bin/bash
# Run this script to deploy code-grader on a Ubuntu machine

export DEBIAN_FRONTEND=noninteractive

sudo apt-get -qq update
sudo apt-get -yqq install docker.io
sudo usermod -aG docker $USER

if [ -z $FROM_SOURCE ]
then
    docker-compose pull -q
else
    docker-compose build -q
fi

# Automated Testing
docker-compose run backend go test -cover ./...
