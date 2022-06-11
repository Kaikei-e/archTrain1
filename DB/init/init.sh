#!/bin/bash 
set -e 
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE DATABASE $POSTGRES_DB WITH OWNER = $POSTGRES_USER;

  GRANT ALL PRIVILEGES ON "public"."users" TO user1; 

  CREATE TABLE "public"."users"(
    "id" integer, PRIMARY KEY ("id"),
    "password" varchar(32) not null,
    "created_at" timestamp not null default current_timestamp,
    "fail_status" smallint not null,
    "token" varchar(32)
  );
EOSQL