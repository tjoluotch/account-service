#!/bin/bash

IMG="account"

run(){
echo "building docker container..."
docker run \
--restart unless-stopped \
-p 700:8080/tcp \
-d \
--env-file config.env \
--network microservice \
--name account-service \
account:0.0.1
}

build(){
echo "building docker image for service..."
docker build \
--compress \
--rm \
-t ${IMG}:0.0.1 .
echo "docker image build for ${IMG} completed..."
}

if [[ $1 == "build" ]]
then
  build
elif [[ $1 == "run" ]]
then
    run
fi
