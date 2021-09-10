#!/bin/bash
# Run this script to deploy code-grader on a Ubuntu machine

# Install Docker
curl -fsSL https://get.docker.com | sudo sh

# Add user to docker group
sudo usermod -aG docker $USER

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
