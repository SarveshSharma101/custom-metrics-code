#!/bin/bash

# Define the Docker repository name and tag
REPO_NAME="sarveshdockerrepo/prometheus"
TAG="service"

# Build the Docker image
docker build -t $REPO_NAME:$TAG .

# Push the Docker image to the repository
docker push $REPO_NAME:$TAG

