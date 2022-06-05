CREATE USER user1;

CREATE DATABASE suiibell;

GRANT ALL PRIVILEGES ON DATABASE suiibell TO user1;

\c user1

CREATE TABLE [IF NOt EXISTS] users (
      id            integer          primary key,
      password      varchar(32)      not null,
      created_at    timestamp        not null default current_timestamp,
      fail_status   smallint         not null,
      token         varchar(32)      ,
      );

      