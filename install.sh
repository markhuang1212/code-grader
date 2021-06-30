#!/bin/bash
# Run this script to deploy code-grader on a Ubuntu machine

USER=$(id -un)

sudo apt-get update
sudo apt-get -y upgrade

sudo apt-get install docker.io
sudo usermod -aG docker 

docker-compose build 
docker-compose up