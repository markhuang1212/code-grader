#!/bin/bash
# Run this script to update code-grader

git pull
docker-compose build
docker-compose restart