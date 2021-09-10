#!/bin/bash
# Run this script to deploy code-grader on a Ubuntu machine

sudo apt-get update
sudo apt-get -y upgrade

sudo apt-get install docker.io
sudo usermod -aG docker $USER

docker-compose build

# Start the server
# docker-compose up
