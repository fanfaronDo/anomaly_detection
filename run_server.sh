#!/bin/bash


if command -v docker-compose > /dev/null; then
    make build_server && make build_client
    sudo docker-compose up -d 
else
    echo "Docker compose does not exist, install and try again"
    exist 1
fi