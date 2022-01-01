drop table if exists Users;
create table Users (
    nickname text unique primary key,
    fullname text not null,
    about text,
    email text not null unique
)

drop table if exists Forum;
create table Forum (
    title text,
    "user" text,
    slug text primary key,
    posts bigint default 0,
    threads bigint default 0,
    foreign key ("user") references users(nickname)
)