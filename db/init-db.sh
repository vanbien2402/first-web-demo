#!/bin/bash

set -e

echo "init database"

PGPASSWORD=$POSTGRESQL_POSTGRES_PASSWORD psql -U postgres -h db-0 <<-EOSQL
CREATE DATABASE common;
EOSQL
PGPASSWORD=$POSTGRES_PASSWORD psql -f ./docker-entrypoint-initdb.d/init/sqls/init-db.sql  -U $POSTGRES_USERNAME -h db-0 -d common