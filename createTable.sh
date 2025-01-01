#!/bin/env bash
# Execute this script after container is up and running
docker cp ./createTable.sql db-for-go:/createTable.sql
docker exec -i db-for-go mysql -udocker -pdocker sampledb < ./createTable.sql

# docker compose exec mysql bash
# mysql -udocker -pdocker sampledb
# select * from articles;
# select * from comments;
# show columns from articles;
# show columns from comments;