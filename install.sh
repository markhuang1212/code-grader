#!/bin/bash
# Run this script to deploy code-grader on a Ubuntu machine

export DEBIAN_FRONTEND=noninteractive

sudo apt-get update
sudo apt-get -y install docker.io
sudo usermod -aG docker $USER

if [ -z $FROM_SOURCE ]
then
    docker-compose pull
else
    docker-compose build
fi

# Automated Testing
docker-compose run backend go test -cover ./...
