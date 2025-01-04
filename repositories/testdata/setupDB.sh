#!/usr/bin/env bash

set -eux

SCRIPT_DIR=$(cd $(dirname $0);pwd)

docker cp $SCRIPT_DIR/setupDB.sql db-for-go:/setupDB.sql
docker exec -i db-for-go mysql -udocker -pdocker sampledb < $SCRIPT_DIR/setupDB.sql