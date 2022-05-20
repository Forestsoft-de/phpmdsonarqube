#!/bin/bash
TAG="v4"

docker build -t forestsoft/phpmdsonarqube:${TAG} .
if [ "$1" == "push" ]; then
    docker push forestsoft/phpmdsonarqube:${TAG}
fi
