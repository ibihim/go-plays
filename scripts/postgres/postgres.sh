#!/usr/bin/env bash

docker run                           \
    -it                              \
    --name test-postgres             \
    -e POSTGRES_USER=ibihim          \
    -e POSTGRES_PASSWORD=helloGithub \
    -e POSTGRES_DB=test              \
    -p 5432:5432                     \
    -d                               \
    postgres
