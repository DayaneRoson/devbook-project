create database if not exists devbook;
use devbook;

drop table if exists followers;
drop table if exists users;


create table users(
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(75) not null,
    createdAt timestamp default current_timestamp()
);

create table followers(
    user_id int not null,
    foreign key (user_id)
    references users(id)
    on delete cascade,

    follower_id int not null,
    foreign key (follower_id)
    references users(id)
    on delete cascade,

    primary key(user_id, follower_id)
);