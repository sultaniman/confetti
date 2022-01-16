CREATE EXTENSION postgis;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id         UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    full_name  VARCHAR(255)                NULL,
    email      VARCHAR(100)                NOT NULL,
    password   VARCHAR(2048)               NOT NULL,
    is_admin   BOOL                        NOT NULL DEFAULT false,
    is_active  BOOL                        NOT NULL DEFAULT false,
    settings   JSONB                       NOT NULL DEFAULT '{}'::jsonb,

    -- 'auth' is default, other authentication providers
    -- for example can be one of google, twitter etc.
    provider   VARCHAR(56)                 NULL     DEFAULT 'auth',
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP)
);

CREATE UNIQUE INDEX ix_users_email ON users (lower(email));
CREATE INDEX ix_users_full_name ON users (lower(full_name));

CREATE TABLE messages
(
    id           UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    owner_id     UUID                        NOT NULL,
    to_user_id   UUID                        NOT NULL,
    message      VARCHAR(1024)               NULL,
    raw_location VARCHAR(32)                 NULL,
    location     GEOMETRY(POINT, 4326)       NULL,
    created_at   TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),
    expires_at   TIMESTAMP WITHOUT TIME ZONE NULL,

    CONSTRAINT fk_messages_owner
        FOREIGN KEY (owner_id)
            REFERENCES users (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_messages_to_user
        FOREIGN KEY (to_user_id)
            REFERENCES users (id)
            ON DELETE CASCADE
);

CREATE INDEX ix_messages_owner ON messages (owner_id);
CREATE INDEX ix_messages_to_user ON messages (to_user_id);
CREATE INDEX ix_messages_geo_location ON messages USING gist (location);

CREATE TABLE files
(
    id         UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    owner_id   UUID NOT NULL,
    message_id UUID NOT NULL,

    CONSTRAINT fk_files_owner
        FOREIGN KEY (owner_id)
            REFERENCES users (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_files_message
        FOREIGN KEY (message_id)
            REFERENCES messages (id)
            ON DELETE CASCADE
);

CREATE INDEX ix_files_owner ON files (owner_id);
CREATE INDEX ix_files_related_message ON files (message_id);

CREATE TABLE friends
(
    id         UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    user_id    UUID                        NOT NULL,
    friend_id  UUID                        NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),

    CONSTRAINT fk_friends_user_id
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_friends_friend_id
        FOREIGN KEY (friend_id)
            REFERENCES users (id)
            ON DELETE CASCADE
);

CREATE TABLE events
(
    id         UUID PRIMARY KEY                     DEFAULT uuid_generate_v4(),
    action     VARCHAR(20)                 NOT NULL, -- create:user, update:user etc.
    context    VARCHAR(20)                 NULL,     -- handler:users service:webhooks etc.
    source     VARCHAR(40)                 NULL,     -- user id, record id etc.
    data       JSONB                       NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc'::text, CURRENT_TIMESTAMP),
    expires_at TIMESTAMP WITHOUT TIME ZONE NULL
);

CREATE INDEX ix_events_action ON events (action);
CREATE INDEX ix_events_context ON events (context);
CREATE INDEX ix_events_source ON events (source);

INSERT INTO users
VALUES (uuid_generate_v4(),
        'Full Admin',
        'admin@example.com',
        'password',
        true,
        true,
        '{}',
        'auth',
        timezone('utc'::text, CURRENT_TIMESTAMP),
        timezone('utc'::text, CURRENT_TIMESTAMP));
