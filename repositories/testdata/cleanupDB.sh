#!/usr/bin/env bash

set -eux

SCRIPT_DIR=$(cd $(dirname $0);pwd)

docker cp $SCRIPT_DIR/cleanupDB.sql db-for-go:/cleanupDB.sql
docker exec -i db-for-go mysql -udocker -pdocker sampledb < $SCRIPT_DIR/cleanupDB.sql