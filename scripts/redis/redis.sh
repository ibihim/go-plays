#!/usr/bin/env bash

CONF_PATH=$(pwd)/auth.conf

docker run                                                      \
    --rm                                                        \
    -it                                                         \
    -p 6379:6379                                                \
    -v "${CONF_PATH}":/usr/local/etc/redis/redis.conf           \
    --name test-redis                                           \
    redis                                                       \
    redis-server /usr/local/etc/redis/redis.conf
