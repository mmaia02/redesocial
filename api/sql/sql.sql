CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varChar(50) not null,
    nick varChar(50) not null unique,
    email varChar(50) not null unique,
    senha varChar(20) not null unique,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;