#!/bin/bash
# Run this script to deploy code-grader on a Ubuntu machine

# Install Docker
if [ -z "$(which docker)" ]
then
    curl -fsSL https://get.docker.com | sudo sh
fi

# Add user to docker group
if [ $USER -ne "root" ]
then
    sudo usermod -aG docker $USER
fi

# Build or pull images
if [ -z $FROM_SOURCE ]
then
    docker-compose pull 
else
    docker-compose build
fi

# Automated Testing
docker-compose run backend go test -cover ./...

# Ready to go!
# docker-compose up
