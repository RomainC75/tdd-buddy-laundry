#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER docker;
    CREATE DATABASE docker;
    GRANT ALL PRIVILEGES ON DATABASE docker TO docker;
    
    CREATE TABLE IF NOT EXISTS reservation (
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