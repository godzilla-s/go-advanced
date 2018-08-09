#!/bin/bash

docker stop $(docker ps -aq)
sleep 0.5
docker rm $(docker ps -aq)
