#!/bin/bash
docker build -t go-server -f go-server/Dockerfile go-server 
docker build -t python-server -f python-server/Dockerfile python-server
docker build -t node-server -f nodejs-server/Dockerfile nodejs-server
