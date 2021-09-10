#!/bin/bash
# Run this script to deploy code-grader on a Ubuntu machine

export DEBIAN_FRONTEND=noninteractive

sudo apt-get -qq update
sudo apt-get -yqq install docker.io
sudo usermod -aG docker $USER

docker-compose build -q

# Automated Testing
docker-compose run backend go test -cover ./...
