#!/usr/bin/env bash
# Execute this script after container is up and running
set -eux

SCRIPT_DIR=$(cd $(dirname $0);pwd)

docker exec -i db-for-go mysql -udocker -pdocker sampledb < $SCRIPT_DIR/createTable.sql

# docker compose exec mysql bash
# mysql -udocker -pdocker sampledb
# select * from articles;
# select * from comments;
# show columns from articles;
# show columns from comments;