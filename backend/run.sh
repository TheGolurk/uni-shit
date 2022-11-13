#!/bin/bash

echo "corriendo el contenedor de base de datos"
docker-compose up -d

echo "compilando la app"
go build -o shituni cmd/main.go

echo "corriendo la app"
./shituni &> out.txt &

#  matar el proceso
# ps -fea
# kill -9 $process-id
# pkill shituni