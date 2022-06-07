CREATE DATABASE suiibell;

USE suiibell;

CREATE SCHEMA user;

GRANT ALL PRIVILEGES ON suiibell.user.users TO user1;

CREATE TABLE users (
      id            integer          primary key,
      password      varchar(32)      not null,
      created_at    timestamp        not null default current_timestamp,
      fail_status   smallint         not null,
      token         varchar(32)      ,
      );

