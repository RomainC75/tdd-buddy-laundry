#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER docker;
    CREATE DATABASE docker;
    GRANT ALL PRIVILEGES ON DATABASE docker TO docker;

    CREATE TABLE reservation (
        id   uuid PRIMARY KEY,
        reservation_date TIMESTAMP NOT NULL,
        reservation_time INT NOT NULL, 
        email VARCHAR NOT NULL,
        pin VARCHAR NOT NULL,
        machine_num VARCHAR NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
    );
EOSQL

# pwd
ls
# for f in /Users/romainchenard/Work/tdd/go-laundry/adapters/secondary/repositories/sqlc/db/migration/*.sql; do
#   echo "Running $f"
#   cat $f
#   echo"----------"
#   psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" -f "$f"
# done