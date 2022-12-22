create database if not exists devbook;
use devbook;

drop table if exists usuarios;

create table users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(50) not null unique,
    createdAt timestamp default current_timestamp()
);