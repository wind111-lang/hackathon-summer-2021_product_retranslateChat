CREATE DATABASE chat_logs;
use chat_logs;

create table chat_logs(
 id          bigint NOT NULL AUTO_INCREMENT,
 name   VARCHAR(256),
 text       longtext,
 time      text,
 PRIMARY KEY (id)
) DEFAULT CHARSET=utf8mb4;