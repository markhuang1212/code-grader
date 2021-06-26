#!/bin/sh

##### Usage: wget -qO - <url> | sudo bash -
##### Requirement: Ubuntu Latest/LTS

PROJ_DIR=/code-grader
RUNTIME_DIR=/code-grader/RUNTIME_DIR
BACKEND_DIR=/code-grader/backend
TEST_CASE_DIR=/code-grader/testcases

SERVICE_FILE=./code-grader.service
SERVICE_DIR=/etc/systemd/system

##### Install Necessary Tools

apt install -y git build-essential

##### Install Docker

##### Install Go

##### Clone Source Code

mkdir $PROJ_DIR
cd $PROJ_DIR
git clone https://github.com/markhuang1212/code-grader .

##### Install & Configure

cd $BACKEND_DIR
make

cd $RUNTIME_DIR
docker build . -t code-grader:0.0.1

cd $PROJ_DIR
mv $SERVICE_FILE $SERVICE_DIR
systemctl enable code-grader