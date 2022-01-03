drop table if exists Users;
create table Users (
    nickname text unique primary key,
    fullname text not null,
    about text,
    email text not null unique
);

drop table if exists Forum;
create table Forum (
    title text,
    "user" text,
    slug text primary key unique,
    posts bigint default 0,
    threads bigint default 0,
    foreign key ("user") references users(nickname)
);

drop table if exists thread;
create table thread (
    id serial primary key,
    title text,
    author text,
    forum text,
    message text,
    votes int default 0,
    slug text unique,
    created timestamp with time zone default now(),
    foreign key (author) references "users"(nickname),
    foreign key (forum)  references "forum" (slug)
);

drop table if exists users_forum;
create table users_forum (
    nickname text not null,
    slug text not null,
    foreign key (nickname) references users(nickname),
    foreign key (slug) references forum (slug),
    unique (nickname, slug)
);

create or replace function new_user_forum() returns trigger as $new_user_forum$
begin
    insert into users_forum (nickname, slug) values (new.author, new.forum) on conflict do nothing;
    return new;
end
$new_user_forum$ language plpgsql;

create trigger insert_thread
    after insert
    on thread
    for each row
    execute procedure new_user_forum();

create or replace function increment_counter_threads() returns trigger as $increment_counter_threads$
begin
    update forum
    set threads = forum.threads + 1
    where slug = new.forum;
    return new;
end
$increment_counter_threads$ language plpgsql;

create trigger increment_counter_threads
    after insert
    on thread
    for each row
execute procedure increment_counter_threads();

drop table if exists post;
create table post(
    id bigserial primary key,
    parent bigint default 0,
    author text,
    message text,
    isEdited boolean default false,
    forum text,
    thread int,
    created timestamp with time zone default now(),
    foreign key (author) references users(nickname),
    foreign key (forum) references forum(slug),
    foreign key (thread) references thread(id)
);

drop table if exists vote;
create table vote (
    id bigserial primary key,
    nickname text,
    voice int,
    thread int not null,
    foreign key (nickname) references users(nickname),
    foreign key (thread) references thread(id),
    unique (thread, nickname)
);