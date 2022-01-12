CREATE EXTENSION IF NOT EXISTS CITEXT;

DROP TABLE IF EXISTS users CASCADE;
create unlogged TABLE users (
                       nickname CITEXT UNIQUE PRIMARY KEY,
                       fullname TEXT NOT NULL,
                       about TEXT,
                       email CITEXT NOT NULL UNIQUE
);


DROP TABLE IF EXISTS forum CASCADE;
create unlogged TABLE forum (
                       title TEXT,
                       "user" CITEXT,
                       slug CITEXT PRIMARY KEY UNIQUE,
                       posts BIGINT DEFAULT 0,
                       threads BIGINT DEFAULT 0,
                       FOREIGN KEY ("user") REFERENCES users(nickname)
);


DROP TABLE IF EXISTS thread CASCADE;
create unlogged TABLE thread (
                        id SERIAL PRIMARY KEY,
                        title TEXT,
                        author CITEXT,
                        forum CITEXT,
                        message TEXT,
                        votes INT DEFAULT 0,
                        slug CITEXT UNIQUE,
                        created TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                        FOREIGN KEY (author) REFERENCES "users"(nickname),
                        FOREIGN KEY (forum)  REFERENCES "forum" (slug)
);


DROP TABLE IF EXISTS users_forum;
CREATE unlogged TABLE users_forum (
                             nickname CITEXT NOT NULL,
                             slug CITEXT NOT NULL,
                             FOREIGN KEY (nickname) REFERENCES users(nickname),
                             FOREIGN KEY (slug) REFERENCES forum (slug),
                             UNIQUE (nickname, slug)
);


DROP TABLE IF EXISTS post CASCADE;
CREATE unlogged TABLE post(
                     id BIGSERIAL PRIMARY KEY,
                     parent BIGINT DEFAULT 0,
                     author CITEXT,
                     message TEXT,
                     isEdited BOOLEAN DEFAULT FALSE,
                     forum CITEXT,
                     thread INT,
                     created TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                     paths BIGINT[] DEFAULT ARRAY []::INTEGER[],
                     FOREIGN KEY (author) REFERENCES users(nickname),
                     FOREIGN KEY (forum) REFERENCES forum(slug),
                     FOREIGN KEY (thread) REFERENCES thread(id)
);


DROP TABLE IF EXISTS vote CASCADE;
CREATE  unlogged TABLE vote (
                      id BIGSERIAL PRIMARY KEY,
                      nickname CITEXT,
                      voice INT,
                      thread INT NOT NULL,
                      FOREIGN KEY (nickname) REFERENCES users(nickname),
                      FOREIGN KEY (thread) REFERENCES thread(id),
                      UNIQUE (thread, nickname)
);


CREATE OR REPLACE FUNCTION add_votes() RETURNS TRIGGER AS
$add_votes$
BEGIN
UPDATE thread SET votes=(votes + NEW.voice) WHERE id = NEW.thread;
return NEW;
end
$add_votes$ LANGUAGE plpgsql;

CREATE TRIGGER after_insert_vote
    AFTER INSERT
    ON vote
    FOR EACH ROW
    EXECUTE PROCEDURE add_votes();


CREATE OR REPLACE FUNCTION update_thread_votes() RETURNS TRIGGER AS
$update_thread_votes$
BEGIN
    IF OLD.voice <> NEW.voice THEN
UPDATE thread SET votes=(votes + NEW.voice * 2) WHERE id = NEW.thread;
END IF;

RETURN NEW;
END
$update_thread_votes$ LANGUAGE plpgsql;

CREATE TRIGGER after_update_voice
    AFTER UPDATE
    ON vote
    FOR EACH ROW
    EXECUTE PROCEDURE update_thread_votes();


CREATE OR REPLACE FUNCTION new_user_forum() RETURNS TRIGGER AS $new_user_forum$
BEGIN
INSERT INTO users_forum (nickname, slug) VALUES (new.author, new.forum) ON CONFLICT DO NOTHING;
RETURN new;
END
$new_user_forum$ LANGUAGE plpgsql;

CREATE TRIGGER after_insert_thread_update_user
    AFTER INSERT
    ON thread
    FOR EACH ROW
    EXECUTE PROCEDURE new_user_forum();

CREATE TRIGGER after_insert_post
    AFTER INSERT
    ON post
    FOR EACH ROW
    EXECUTE PROCEDURE new_user_forum();


CREATE OR REPLACE FUNCTION update_paths_post() RETURNS TRIGGER AS
$update_paths_post$
DECLARE
parent_path         BIGINT[];
    first_parent_thread INT;
BEGIN
    IF (NEW.parent = 0) THEN
        NEW.paths := array_append(NEW.paths, NEW.id);
ELSE
SELECT paths FROM post WHERE id = NEW.parent INTO parent_path;
SELECT thread FROM post WHERE id = parent_path[1] INTO first_parent_thread;

IF NOT FOUND OR first_parent_thread <> NEW.thread THEN
            RAISE EXCEPTION 'parent post was created in another thread' USING ERRCODE = '77777';
END IF;

        NEW.paths := NEW.paths || parent_path || NEW.id;
END IF;

UPDATE forum SET posts=posts + 1 WHERE forum.slug = NEW.forum;

RETURN NEW;
END
$update_paths_post$ LANGUAGE plpgsql;

CREATE TRIGGER before_insert_post
    BEFORE INSERT
    ON post
    FOR EACH ROW
    EXECUTE PROCEDURE update_paths_post();


CREATE OR REPLACE FUNCTION increment_counter_threads() RETURNS TRIGGER AS $increment_counter_threads$
BEGIN
UPDATE forum
SET threads = forum.threads + 1
WHERE slug = NEW.forum;
RETURN NEW;
END
$increment_counter_threads$ LANGUAGE plpgsql;

CREATE TRIGGER after_insert_thread
    AFTER INSERT
    ON thread
    FOR EACH ROW
    EXECUTE PROCEDURE increment_counter_threads();

create index if not exists users_nickname on users using hash(nickname);
create index if not exists users_email on users using hash(email);

create index if not exists post_thread on post(thread);
create index if not exists post_null_parent on post(id, thread) where parent = 0;
create index if not exists post_thread_paths_asc on post(thread, paths);
create index if not exists post_thread_path_path on post(thread, (paths[0]), paths);
create index if not exists post_thread_created_id on post(thread, created, id);

create index if not exists forum_slug on forum using hash(slug);

create index if not exists thread_forum on thread(forum);
create index if not exists thread_forum_created on thread(forum, created);
create unique index if not exists thread_slug on thread (slug);

