drop table if exists Users;
create table Users (
    nickname text unique primary key,
    fullname text not null,
    about text,
    email text not null unique
)